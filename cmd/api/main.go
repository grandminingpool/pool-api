package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-playground/validator/v10"
	appConfig "github.com/grandminingpool/pool-api/configs/app"
	postgresConfig "github.com/grandminingpool/pool-api/configs/postgres"
	apiServer "github.com/grandminingpool/pool-api/internal/api/server"
	blocksServices "github.com/grandminingpool/pool-api/internal/api/services/blocks"
	chartsServices "github.com/grandminingpool/pool-api/internal/api/services/charts"
	minersServices "github.com/grandminingpool/pool-api/internal/api/services/miners"
	payoutsServices "github.com/grandminingpool/pool-api/internal/api/services/payouts"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/flags"
	"github.com/grandminingpool/pool-api/internal/common/logger"
	postgresProvider "github.com/grandminingpool/pool-api/internal/providers/postgres"
	"go.uber.org/zap"
)

func main() {
	//	Init context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	//	Parse flags
	parsedFlags := flags.ParseFlags()

	//	Setup flags
	flagsConf := flags.SetupFlags(parsedFlags)

	//	Setup logger
	zapLogger, err := logger.SetupLogger(&logger.LoggerConfig{
		AppMode:         flagsConf.Mode,
		OutputPath:      flagsConf.Logger.OutputPath,
		ErrorOutputPath: flagsConf.Logger.ErrorOutputPath,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to setup zap logger: %w", err))
	}
	defer zapLogger.Sync()

	zap.ReplaceGlobals(zapLogger)

	//	Init validator
	validate := validator.New()

	//	Init postgres config
	postgresConf, err := postgresConfig.New(flagsConf.ConfigsPath, validate)
	if err != nil {
		zap.L().Fatal("failed to load postgres config", zap.Error(err))
	}

	//	Init postgres connection
	pgConn, err := postgresProvider.NewConnection(ctx, postgresConf)
	if err != nil {
		zap.L().Fatal("failed to create postgres connection", zap.Error(err))
	}

	zap.L().Info("successfully connected to postgres database")

	//	Init app config
	appConf, err := appConfig.New(flagsConf.ConfigsPath, validate)
	if err != nil {
		zap.L().Fatal("failed to load application config", zap.Error(err))
	}

	//	Init blockchains service and start
	blockchainsService := blockchains.NewService(pgConn, &appConf.PoolAPI)
	if err := blockchainsService.Start(ctx, flagsConf.CertsPath); err != nil {
		zap.L().Fatal("failed to start blockchains service", zap.Error(err))
	}

	//	Init api server
	poolsBlockchainService := &poolsServices.BlockchainService{}
	pricesService := pricesServices.NewPricesService(pgConn)
	minersBlockchainService := &minersServices.BlockchainService{}
	payoutsBlockchainService := &payoutsServices.BlockchainService{}
	blocksBlockchainService := &blocksServices.BlockchainService{}
	chartsBlockchainService := &chartsServices.BlockchainService{}
	apiHandlers := apiServer.CreateHandler(
		blockchainsService,
		poolsBlockchainService,
		pricesService,
		minersBlockchainService,
		payoutsBlockchainService,
		blocksBlockchainService,
		chartsBlockchainService,
	)
	apiSrv, err := apiServer.CreateServer(apiHandlers)
	if err != nil {
		zap.L().Fatal("failed to create pool api server", zap.Error(err))
	}

	apiHTTPServer := &http.Server{Addr: appConf.Address(), Handler: apiSrv}

	//	Subscribe to system signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		stop := <-signalChan

		zap.L().Info("waiting for all processes to stop", zap.String("signal", stop.String()))

		var stopErr error
		if stopErr = apiHTTPServer.Shutdown(ctx); stopErr != nil {
			zap.L().Fatal("failed to shutdown pool api server", zap.Error(err))
		}

		cancel()

		blockchainsService.Close()
		zap.L().Info("closed blockchains pool api connections")

		if stopErr = pgConn.Close(); stopErr != nil {
			zap.L().Fatal("failed to close postgres connection", zap.Error(stopErr))
		}

		zap.L().Info("closed postgres connection")
	}()

	//	Run the server
	zap.L().Info(fmt.Sprintf("starting pool api http server on: %s", appConf.Address()))

	if err = apiHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("failed to start listen pool api http server address", zap.Error(err))
	}

	wg.Wait()
	zap.L().Info("pool api stopped")
}
