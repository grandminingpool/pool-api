// Code generated by ogen, DO NOT EDIT.

package apiModels

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetBlockchainPool implements getBlockchainPool operation.
	//
	// Get full pool data on blockchain.
	//
	// GET /pools/{blockchain}
	GetBlockchainPool(ctx context.Context, params GetBlockchainPoolParams) (*Pool, error)
	// GetBlockchainPoolInfo implements getBlockchainPoolInfo operation.
	//
	// Get pool info on blockchain.
	//
	// GET /pools/{blockchain}/info
	GetBlockchainPoolInfo(ctx context.Context, params GetBlockchainPoolInfoParams) (*PoolInfo, error)
	// GetBlockchainPoolSlaves implements getBlockchainPoolSlaves operation.
	//
	// Get pool locations on blockchain.
	//
	// GET /pools/{blockchain}/slaves
	GetBlockchainPoolSlaves(ctx context.Context, params GetBlockchainPoolSlavesParams) ([]PoolSlave, error)
	// GetBlockchainPoolStats implements getBlockchainPoolStats operation.
	//
	// Get pool statistics on blockchain.
	//
	// GET /pools/{blockchain}/stats
	GetBlockchainPoolStats(ctx context.Context, params GetBlockchainPoolStatsParams) (PoolStats, error)
	// GetBlockchainPrice implements getBlockchainPrice operation.
	//
	// Get blockchain coin price and markets.
	//
	// GET /prices/{blockchain}
	GetBlockchainPrice(ctx context.Context, params GetBlockchainPriceParams) (*BlockchainCoinPrice, error)
	// GetBlockchains implements getBlockchains operation.
	//
	// Get available blockchains list.
	//
	// GET /blockchains
	GetBlockchains(ctx context.Context) ([]Blockchain, error)
	// GetPrices implements getPrices operation.
	//
	// Get pool blockchain coin price list.
	//
	// GET /prices
	GetPrices(ctx context.Context) ([]CoinPrice, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}