package apiServer

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blockchainsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blockchains"
	blocksHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blocks"
	chartsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/charts"
	minersHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/miners"
	payoutsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/payouts"
	poolsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/pools"
	pricesHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/prices"
	blockchainsSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/blockchains"
	blocksSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/blocks"
	chartsSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/charts"
	minersSerializer "github.com/grandminingpool/pool-api/internal/api/serializers/miners"
	payoutsSerializer "github.com/grandminingpool/pool-api/internal/api/serializers/payouts"
	poolsSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/pools"
	pricesSerializer "github.com/grandminingpool/pool-api/internal/api/serializers/prices"
	apiServerHandlers "github.com/grandminingpool/pool-api/internal/api/server/handlers"
	blocksServices "github.com/grandminingpool/pool-api/internal/api/services/blocks"
	chartsServices "github.com/grandminingpool/pool-api/internal/api/services/charts"
	minersServices "github.com/grandminingpool/pool-api/internal/api/services/miners"
	payoutsServices "github.com/grandminingpool/pool-api/internal/api/services/payouts"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

func CreateHandler(
	blockchainsService *blockchains.Service,
	poolsBlockchainService *poolsServices.BlockchainService,
	poolsService *poolsServices.PoolsService,
	pricesService *pricesServices.PricesService,
	minersBlockchainService *minersServices.BlockchainService,
	payoutsBlockchainService *payoutsServices.BlockchainService,
	blocksBlockchainService *blocksServices.BlockchainService,
	chartsBlockchainService *chartsServices.BlockchainService,
) apiModels.Handler {
	//	Init blockchains handlers
	blockchainInfoSerializer := &blockchainsSerializers.BlockchainInfoSerializer{}
	blockchainsHandler := blockchainsHandlers.NewHandler(blockchainsService, blockchainInfoSerializer)

	//	Init pools handlers
	poolInfoSerializer := &poolsSerializers.PoolInfoSerializer{}
	poolStatsSerializer := &poolsSerializers.PoolStatsSerializer{}
	poolSlaveSerializer := &poolsSerializers.PoolSlaveSerialzier{}
	poolSerializer := poolsSerializers.NewPoolSerializer(poolInfoSerializer, poolStatsSerializer, poolSlaveSerializer)
	poolsBlockchainHandler := poolsHandlers.NewBlockchainHandler(
		poolsBlockchainService,
		poolSerializer,
		poolInfoSerializer,
		poolStatsSerializer,
		poolSlaveSerializer,
	)
	blockchainPoolStatsSerializer := poolsSerializers.NewBlockchainPoolStatsSerializer(poolStatsSerializer)
	poolsHandler := poolsHandlers.NewHandler(poolsService, blockchainPoolStatsSerializer)

	//	Init prices handlers
	marketPriceSerializer := &pricesSerializer.MarkerPriceSerializer{}
	blockchainMarketsSerializer := pricesSerializer.NewBlockchainMarketsSerializer(marketPriceSerializer)
	pricesBlockchainHandler := pricesHandlers.NewBlockchainHandler(pricesService, blockchainMarketsSerializer)
	blockchainPriceSerializer := &pricesSerializer.BlockchainPriceSerializer{}
	pricesHandler := pricesHandlers.NewHandler(pricesService, blockchainPriceSerializer)

	//	Init miners handlers
	minerSerializer := &minersSerializer.MinerSerializer{}
	minerWorkerSerializer := &minersSerializer.MinerWorkerSerializer{}
	minersBlockchainHandler := minersHandlers.NewBlockchainHandler(minersBlockchainService, minerSerializer, minerWorkerSerializer)

	//	Init payouts handlers
	payoutSerializer := &payoutsSerializer.PayoutSerializer{}
	payoutsBlockchainHandler := payoutsHandlers.NewBlockchainHandler(payoutsBlockchainService, payoutSerializer)

	//	Init blocks handlers
	minedBlockSerializer := &blocksSerializers.MinedBlockSerializer{}
	minedSoloBlockSerializer := &blocksSerializers.MinedSoloBlockSerializer{}
	blocksBlockchainHandler := blocksHandlers.NewBlockchainHandler(blocksBlockchainService, minedBlockSerializer, minedSoloBlockSerializer)

	//	Init charts handlers
	poolStatsPointSerializer := &chartsSerializers.PoolStatsPointSerializer{}
	poolDifficultiesPointSerializer := &chartsSerializers.PoolDifficultiesPointSerializer{}
	roundsPointSerializer := &chartsSerializers.RoundsPointSerializer{}
	minerHashratesPointSerializer := &chartsSerializers.MinerHashratesPointSerializer{}
	minerSharesPointSerializer := &chartsSerializers.MinerSharesPointSerializer{}
	chartsBlockchainHandler := chartsHandlers.NewBlockchainHandler(
		chartsBlockchainService,
		poolStatsPointSerializer,
		poolDifficultiesPointSerializer,
		roundsPointSerializer,
		minerHashratesPointSerializer,
		minerSharesPointSerializer,
	)

	return apiServerHandlers.NewServerHandler(
		blockchainsHandler,
		poolsBlockchainHandler,
		poolsHandler,
		pricesBlockchainHandler,
		pricesHandler,
		minersBlockchainHandler,
		payoutsBlockchainHandler,
		blocksBlockchainHandler,
		chartsBlockchainHandler,
		blockchainsService,
	)
}
