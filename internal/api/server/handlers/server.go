package apiServerHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blockchainsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blockchains"
	poolsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/pools"
	pricesHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/prices"
	apiServerErrors "github.com/grandminingpool/pool-api/internal/api/server/errors"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

type ServerHandler struct {
	blockchainsHandler      *blockchainsHandlers.BlockchainsHandler
	poolsBlockchainHandler  *poolsHandlers.BlockchainHandler
	pricesBlockchainHandler *pricesHandlers.BlockchainHandler
	blockchainService       *blockchains.Service
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
