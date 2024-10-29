package apiServerHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blockchainsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blockchains"
	blocksHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blocks"
	blocksQuery "github.com/grandminingpool/pool-api/internal/api/handlers/blocks/query"
	minersHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/miners"
	minersQuery "github.com/grandminingpool/pool-api/internal/api/handlers/miners/query"
	payoutsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/payouts"
	payoutsQuery "github.com/grandminingpool/pool-api/internal/api/handlers/payouts/query"
	poolsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/pools"
	pricesHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/prices"
	apiServerErrors "github.com/grandminingpool/pool-api/internal/api/server/errors"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

type ServerHandler struct {
	blockchainsHandler       *blockchainsHandlers.Handler
	poolsBlockchainHandler   *poolsHandlers.BlockchainHandler
	pricesBlockchainHandler  *pricesHandlers.BlockchainHandler
	pricesHandler            *pricesHandlers.Handler
	minersBlockchainHandler  *minersHandlers.BlockchainHandler
	payoutsBlockchainHandler *payoutsHandlers.BlockchainHandler
	blocksBlockchainHandler  *blocksHandlers.BlockchainHandler
	blockchainService        *blockchains.Service
}

func (h *ServerHandler) GetBlockchains(ctx context.Context) ([]apiModels.Blockchain, error) {
	return h.blockchainsHandler.Get(ctx)
}

func (h *ServerHandler) getBlockchain(coin string) (*blockchains.Blockchain, error) {
	blockchain, err := h.blockchainService.GetBlockchain(coin)
	if err != nil {
		return nil, serverErrors.CreateNotFoundError(apiServerErrors.BlockchainNotFound, err)
	}

	return blockchain, nil
}

func (h *ServerHandler) GetBlockchainPool(ctx context.Context, params apiModels.GetBlockchainPoolParams) (*apiModels.Pool, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.poolsBlockchainHandler.GetPool(ctx, blockchain)
}

func (h *ServerHandler) GetBlockchainPoolInfo(ctx context.Context, params apiModels.GetBlockchainPoolInfoParams) (*apiModels.PoolInfo, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.poolsBlockchainHandler.GetPoolInfo(ctx, blockchain)
}

func (h *ServerHandler) GetBlockchainPoolStats(ctx context.Context, params apiModels.GetBlockchainPoolStatsParams) (*apiModels.PoolStats, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.poolsBlockchainHandler.GetPoolStats(ctx, blockchain)
}

func (h *ServerHandler) GetBlockchainPoolSlaves(ctx context.Context, params apiModels.GetBlockchainPoolStatsParams) ([]apiModels.PoolSlave, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.poolsBlockchainHandler.GetPoolSlaves(ctx, blockchain)
}

func (h *ServerHandler) GetBlockchainPrice(ctx context.Context, params apiModels.GetBlockchainPriceParams) (*apiModels.BlockchainCoinPrice, error) {
	return h.pricesBlockchainHandler.GetPrice(ctx, params.Blockchain)
}

func (h *ServerHandler) GetPrices(ctx context.Context) ([]apiModels.CoinPrice, error) {
	return h.pricesHandler.Get(ctx)
}

func (h *ServerHandler) GetBlockchainMiners(ctx context.Context, params apiModels.GetBlockchainMinersParams) (*apiModels.MinersList, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	sorts := minersQuery.ParseMinersSortsInQuery(&params.Sorts)
	filters := minersQuery.ParseMinersFiltersInQueryParams(&params)

	return h.minersBlockchainHandler.GetMiners(ctx, blockchain, sorts, filters, params.Limit, params.Offset)
}

func (h *ServerHandler) GetBlockchainMiner(ctx context.Context, params apiModels.GetBlockchainMinerParams) (*apiModels.Miner, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.minersBlockchainHandler.GetMiner(ctx, blockchain, params.Miner)
}

func (h *ServerHandler) GetBlockchainMinerWorkers(ctx context.Context, params apiModels.GetBlockchainMinerWorkersParams) ([]apiModels.MinerWorker, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.minersBlockchainHandler.GetMinerWorkers(ctx, blockchain, params.Miner)
}

func (h *ServerHandler) GetBlockchainPayouts(ctx context.Context, params apiModels.GetBlockchainPayoutsParams) (*apiModels.PayoutsList, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	sorts := payoutsQuery.ParsePayoutsSortsInQuery(&params.Sorts)
	filters := payoutsQuery.ParsePayoutsFiltersInQueryParams(&params)

	return h.payoutsBlockchainHandler.GetPayouts(ctx, blockchain, sorts, filters, params.Limit, params.Offset)
}

func (h *ServerHandler) GetBlockchainMinerBalance(ctx context.Context, params apiModels.GetBlockchainMinerBalanceParams) (*apiModels.MinerBalance, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	return h.payoutsBlockchainHandler.GetMinerBalance(ctx, blockchain, params.Miner)
}

func (h *ServerHandler) GetBlockchainBlocks(ctx context.Context, params apiModels.GetBlockchainBlocksParams) (*apiModels.MinedBlocksList, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	sorts := blocksQuery.ParseBlocksSortsInQuery(&params.Sorts)
	filters := blocksQuery.ParseBlocksFiltersInQueryParams(&params)

	return h.blocksBlockchainHandler.GetBlocks(ctx, blockchain, sorts, filters, params.Limit, params.Offset)
}

func (h *ServerHandler) GetBlockchainSoloBlocks(ctx context.Context, params apiModels.GetBlockchainSoloBlocksParams) (apiModels.GetBlockchainSoloBlocksRes, error) {
	blockchain, err := h.getBlockchain(params.Blockchain)
	if err != nil {
		return nil, err
	}

	sorts := blocksQuery.ParseSoloBlocksSortsInQuery(&params.Sorts)
	filters := blocksQuery.ParseSoloBlocksFiltersInQueryParams(&params)

	return h.blocksBlockchainHandler.GetSoloBlocks(ctx, blockchain, sorts, filters, params.Limit, params.Offset)
}
