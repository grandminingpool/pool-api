package payoutsServices

import (
	"context"
	"fmt"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	paginationProto "github.com/grandminingpool/pool-api-proto/generated/utils/pagination"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainService struct{}

func (s *BlockchainService) GetPayouts(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.PayoutsSorts,
	filters *poolPayoutsProto.PayoutsFilters,
	limit, offset uint32,
) (*poolPayoutsProto.PayoutsList, error) {
	client := poolPayoutsProto.NewPoolPayoutsServiceClient(blockchain.GetConnection())
	payoutsList, err := client.GetPayouts(ctx, &poolPayoutsProto.GetPayoutsRequest{
		Filters: filters,
		Sorts:   sorts,
		Pagination: &paginationProto.PaginationRequest{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) payouts: %w", blockchain.GetInfo().Coin, err)
	}

	return payoutsList, nil
}

func (s *BlockchainService) GetMinerBalance(ctx context.Context, blockchain *blockchains.Blockchain, miner string) (*uint64, error) {
	client := poolPayoutsProto.NewPoolPayoutsServiceClient(blockchain.GetConnection())
	balancesMap, err := client.GetMinersBalancesFromList(ctx, &poolMinersProto.MinerAddressesRequest{
		Addresses: []string{miner},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) miner (address: %s) balance: %w", blockchain.GetInfo().Coin, miner, err)
	}

	minerBalance, ok := balancesMap.Balances[miner]
	if !ok {
		return nil, nil
	}

	return &minerBalance.Balance, nil
}
