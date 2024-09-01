package apiResolvers

import (
	"context"

	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type HandlerWithBlockchain[T any] func(ctx context.Context, blockchain *blockchains.Blockchain) (*T, error)

type BlockchainResolver[T any] struct {
	blockchainService *blockchains.Service
}

func (r *BlockchainResolver[T]) Handler(
	ctx context.Context,
	blockchainCoin string,
	handler HandlerWithBlockchain[T],
) (*T, error) {
	blockchain, err := r.blockchainService.GetBlockchain(blockchainCoin)
	if err != nil {
		return nil, err
	}

	return handler(ctx, blockchain)
}
