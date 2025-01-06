package apiServerHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blockchainsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blockchains"
	blocksHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blocks"
	blocksQuery "github.com/grandminingpool/pool-api/internal/api/handlers/blocks/query"
	chartsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/charts"
	minersHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/miners"
	minersQuery "github.com/grandminingpool/pool-api/internal/api/handlers/miners/query"
	payoutsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/payouts"
	payoutsQuery "github.com/grandminingpool/pool-api/internal/api/handlers/payouts/query"
	poolsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/pools"
	pricesHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/prices"
	apiServerErrors "github.com/grandminingpool/pool-api/internal/api/server/errors"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type ServerHandler struct {
	blockchainsHandler       *blockchainsHandlers.Handler
	poolsBlockchainHandler   *poolsHandlers.BlockchainHandler
	poolsHandler             *poolsHandlers.Handler
	pricesBlockchainHandler  *pricesHandlers.BlockchainHandler
	pricesHandler            *pricesHandlers.Handler
	minersBlockchainHandler  *minersHandlers.BlockchainHandler
	payoutsBlockchainHandler *payoutsHandlers.BlockchainHandler
	blocksBlockchainHandler  *blocksHandlers.BlockchainHandler
	chartsBlockchainHandler  *chartsHandlers.BlockchainHandler
	blockchainService        *blockchains.Service
}

func (h *ServerHandler) GetBlockchains(ctx context.Context) (*apiModels.BlockchainsList, error) {
	return h.blockchainsHandler.Get(ctx), nil
}

func (h *ServerHandler) getBlockchain(coin string) (*blockchains.Blockchain, *apiModels.BlockchainNotFound) {
	blockchain, err := h.blockchainService.GetBlockchain(coin)
	if err != nil {
		return nil, &apiModels.BlockchainNotFound{
			Code:    string(apiServerErrors.BlockchainNotFoundError),
			Message: err.Error(),
		}
	}

	return blockchain, nil
}

func (h *ServerHandler) GetPools(ctx context.Context, params apiModels.GetPoolsParams) (apiModels.GetPoolsRes, error) {
	return h.poolsHandler.GetPools(ctx, params.IncludeSoloStats.Value, params.IncludeNetworkInfo.Value), nil
}

func (h *ServerHandler) GetBlockchainPool(ctx context.Context, params apiModels.GetBlockchainPoolParams) (apiModels.GetBlockchainPoolRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.poolsBlockchainHandler.GetPool(ctx, blockchain, params.Solo.Value), nil
}

func (h *ServerHandler) GetBlockchainPoolInfo(ctx context.Context, params apiModels.GetBlockchainPoolInfoParams) (apiModels.GetBlockchainPoolInfoRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.poolsBlockchainHandler.GetPoolInfo(ctx, blockchain), nil
}

func (h *ServerHandler) GetBlockchainPoolStats(ctx context.Context, params apiModels.GetBlockchainPoolStatsParams) (apiModels.GetBlockchainPoolStatsRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.poolsBlockchainHandler.GetPoolStats(ctx, blockchain, params.Solo.Value), nil
}

func (h *ServerHandler) GetBlockchainPoolNetworkInfo(ctx context.Context, params apiModels.GetBlockchainPoolNetworkInfoParams) (apiModels.GetBlockchainPoolNetworkInfoRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.poolsBlockchainHandler.GetPoolNetworkInfo(ctx, blockchain), nil
}

func (h *ServerHandler) GetBlockchainPoolSlaves(ctx context.Context, params apiModels.GetBlockchainPoolSlavesParams) (apiModels.GetBlockchainPoolSlavesRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.poolsBlockchainHandler.GetPoolSlaves(ctx, blockchain, params.Solo.Value), nil
}

func (h *ServerHandler) GetBlockchainMarkets(ctx context.Context, params apiModels.GetBlockchainMarketsParams) (apiModels.GetBlockchainMarketsRes, error) {
	return h.pricesBlockchainHandler.GetMarkets(ctx, params.Blockchain), nil
}

func (h *ServerHandler) GetPrices(ctx context.Context) (apiModels.GetPricesRes, error) {
	return h.pricesHandler.Get(ctx), nil
}

func (h *ServerHandler) GetBlockchainMiners(ctx context.Context, params apiModels.GetBlockchainMinersParams) (apiModels.GetBlockchainMinersRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	sorts := minersQuery.ParseMinersSortsInQuery(&params.Sorts)
	filters := minersQuery.ParseMinersFiltersInQueryParams(&params)

	return h.minersBlockchainHandler.GetMiners(ctx, blockchain, sorts, filters, params.Limit, params.Offset), nil
}

func (h *ServerHandler) GetBlockchainMiner(ctx context.Context, params apiModels.GetBlockchainMinerParams) (apiModels.GetBlockchainMinerRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.minersBlockchainHandler.GetMiner(ctx, blockchain, params.Miner), nil
}

func (h *ServerHandler) GetBlockchainMinerWorkers(ctx context.Context, params apiModels.GetBlockchainMinerWorkersParams) (apiModels.GetBlockchainMinerWorkersRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.minersBlockchainHandler.GetMinerWorkers(ctx, blockchain, params.Miner), nil
}

func (h *ServerHandler) GetBlockchainPayouts(ctx context.Context, params apiModels.GetBlockchainPayoutsParams) (apiModels.GetBlockchainPayoutsRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	sorts := payoutsQuery.ParsePayoutsSortsInQuery(&params.Sorts)
	filters := payoutsQuery.ParsePayoutsFiltersInQueryParams(&params)

	return h.payoutsBlockchainHandler.GetPayouts(ctx, blockchain, sorts, filters, params.Limit, params.Offset), nil
}

func (h *ServerHandler) GetBlockchainMinerBalance(ctx context.Context, params apiModels.GetBlockchainMinerBalanceParams) (apiModels.GetBlockchainMinerBalanceRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.payoutsBlockchainHandler.GetMinerBalance(ctx, blockchain, params.Miner), nil
}

func (h *ServerHandler) GetBlockchainBlocks(ctx context.Context, params apiModels.GetBlockchainBlocksParams) (apiModels.GetBlockchainBlocksRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	sorts := blocksQuery.ParseBlocksSortsInQuery(&params.Sorts)
	filters := blocksQuery.ParseBlocksFiltersInQueryParams(&params)

	return h.blocksBlockchainHandler.GetBlocks(ctx, blockchain, sorts, filters, params.Limit, params.Offset), nil
}

func (h *ServerHandler) GetBlockchainSoloBlocks(ctx context.Context, params apiModels.GetBlockchainSoloBlocksParams) (apiModels.GetBlockchainSoloBlocksRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	sorts := blocksQuery.ParseSoloBlocksSortsInQuery(&params.Sorts)
	filters := blocksQuery.ParseSoloBlocksFiltersInQueryParams(&params)

	return h.blocksBlockchainHandler.GetSoloBlocks(ctx, blockchain, sorts, filters, params.Limit, params.Offset), nil
}

func (h *ServerHandler) GetBlockchainPoolStatsChart(ctx context.Context, params apiModels.GetBlockchainPoolStatsChartParams) (apiModels.GetBlockchainPoolStatsChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetPoolStatsChart(ctx, blockchain, &params.Period, &params.Solo), nil
}


func (h *ServerHandler) GetBlockchainRoundsChart(ctx context.Context, params apiModels.GetBlockchainRoundsChartParams) (apiModels.GetBlockchainRoundsChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetRoundsChart(ctx, blockchain, &params.Period), nil
}

func (h *ServerHandler) GetBlockchainMinerHashratesChart(ctx context.Context, params apiModels.GetBlockchainMinerHashratesChartParams) (apiModels.GetBlockchainMinerHashratesChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetMinerHashratesChart(ctx, blockchain, &params.Period, params.Miner, &params.Solo), nil
}

func (h *ServerHandler) GetBlockchainMinerSharesChart(ctx context.Context, params apiModels.GetBlockchainMinerSharesChartParams) (apiModels.GetBlockchainMinerSharesChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetMinerSharesChart(ctx, blockchain, &params.Period, params.Miner, &params.Solo), nil
}

func (h *ServerHandler) GetBlockchainMinerWorkerHashratesChart(ctx context.Context, params apiModels.GetBlockchainMinerWorkerHashratesChartParams) (apiModels.GetBlockchainMinerWorkerHashratesChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetMinerWorkerHashratesChart(ctx, blockchain, &params.Period, params.Miner, params.Worker), nil
}

func (h *ServerHandler) GetBlockchainMinerWorkerSharesChart(ctx context.Context, params apiModels.GetBlockchainMinerWorkerSharesChartParams) (apiModels.GetBlockchainMinerWorkerSharesChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetMinerWorkerSharesChart(ctx, blockchain, &params.Period, params.Miner, params.Worker), nil
}

func (h *ServerHandler) GetBlockchainMinerProfitabilitiesChart(ctx context.Context, params apiModels.GetBlockchainMinerProfitabilitiesChartParams) (apiModels.GetBlockchainMinerProfitabilitiesChartRes, error) {
	blockchain, notFound := h.getBlockchain(params.Blockchain)
	if notFound != nil {
		return notFound, nil
	}

	return h.chartsBlockchainHandler.GetMinerProfitabilitiesChart(ctx, blockchain, &params.Period, params.Miner, params.Solo.Value), nil
}

func NewServerHandler(
	blockchainsHandler *blockchainsHandlers.Handler,
	poolsBlockchainHandler *poolsHandlers.BlockchainHandler,
	poolsHandler *poolsHandlers.Handler,
	pricesBlockchainHandler *pricesHandlers.BlockchainHandler,
	pricesHandler *pricesHandlers.Handler,
	minersBlockchainHandler *minersHandlers.BlockchainHandler,
	payoutsBlockchainHandler *payoutsHandlers.BlockchainHandler,
	blocksBlockchainHandler *blocksHandlers.BlockchainHandler,
	chartsBlockchainHandler *chartsHandlers.BlockchainHandler,
	blockchainService *blockchains.Service,
) *ServerHandler {
	return &ServerHandler{
		blockchainsHandler:       blockchainsHandler,
		poolsBlockchainHandler:   poolsBlockchainHandler,
		poolsHandler:             poolsHandler,
		pricesBlockchainHandler:  pricesBlockchainHandler,
		pricesHandler:            pricesHandler,
		minersBlockchainHandler:  minersBlockchainHandler,
		payoutsBlockchainHandler: payoutsBlockchainHandler,
		blocksBlockchainHandler:  blocksBlockchainHandler,
		chartsBlockchainHandler:  chartsBlockchainHandler,
		blockchainService:        blockchainService,
	}
}
