package minersServices

import (
	"context"
	"fmt"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	paginationProto "github.com/grandminingpool/pool-api-proto/generated/utils/pagination"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainService struct{}

func (s *BlockchainService) GetMiners(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolMinersProto.MinersSorts,
	filters *poolMinersProto.MinersFilters,
	limit, offset uint32,
) (*poolMinersProto.MinersList, error) {
	client := poolMinersProto.NewPoolMinersServiceClient(blockchain.GetConnection())
	minersList, err := client.GetMiners(ctx, &poolMinersProto.GetMinersRequest{
		Filters: filters,
		Sorts:   sorts,
		Pagination: &paginationProto.PaginationRequest{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) miners: %w", blockchain.GetInfo().Coin, err)
	}

	return minersList, nil
}

func (s *BlockchainService) GetMiner(ctx context.Context, blockchain *blockchains.Blockchain, miner string) (*poolMinersProto.Miner, error) {
	client := poolMinersProto.NewPoolMinersServiceClient(blockchain.GetConnection())
	minersList, err := client.GetMiners(ctx, &poolMinersProto.GetMinersRequest{
		Filters: &poolMinersProto.MinersFilters{
			Address: &miner,
		},
		Pagination: &paginationProto.PaginationRequest{
			Limit:  1,
			Offset: 0,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) miner (address: %s), error: %w", blockchain.GetInfo().Coin, miner, err)
	} else if len(minersList.Miners) == 0 {
		return nil, nil
	}

	return minersList.Miners[0], nil
}

func (s *BlockchainService) GetMinerWorkers(ctx context.Context, blockchain *blockchains.Blockchain, miner string) ([]*poolMinersProto.MinerWorker, error) {
	client := poolMinersProto.NewPoolMinersServiceClient(blockchain.GetConnection())
	minerWorkers, err := client.GetMinerWorkers(ctx, &poolMinersProto.MinerAddressRequest{
		Address: miner,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) miner (address: %s) workers: %w", blockchain.GetInfo().Coin, miner, err)
	}

	return minerWorkers.Workers, nil
}
