package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	postgresConfig "github.com/grandminingpool/pool-api/configs/postgres"
	coinPrices "github.com/grandminingpool/pool-api/internal/coin_prices"
	"github.com/grandminingpool/pool-api/internal/common/flags"
	"github.com/grandminingpool/pool-api/internal/common/logger"
	postgresProvider "github.com/grandminingpool/pool-api/internal/providers/postgres"
	"go.uber.org/zap"
)

func main() {
	//	Init context from background
	ctx := context.Background()

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

	//	Init update prices service and start it
	updateService := coinPrices.NewUpdateService(pgConn)
	if err := updateService.Update(ctx); err != nil {
		zap.L().Fatal("failed to update prices", zap.Error(err))
	}

	zap.L().Info("prices was successfully updated")

	//	Close postgres connection
	if err := pgConn.Close(); err != nil {
		zap.L().Fatal("failed to close postgres connection", zap.Error(err))
	}

	zap.L().Info("closed postgres connection")
}
