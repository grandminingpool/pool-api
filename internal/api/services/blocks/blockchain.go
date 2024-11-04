package blocksServices

import (
	"context"
	"fmt"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	paginationProto "github.com/grandminingpool/pool-api-proto/generated/utils/pagination"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BlockchainService struct{}

func (s *BlockchainService) getProtoClient(blockchain *blockchains.Blockchain) poolPayoutsProto.PoolPayoutsServiceClient {
	return poolPayoutsProto.NewPoolPayoutsServiceClient(blockchain.GetConnection())
}

func (s *BlockchainService) GetBlocks(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.MinedBlocksSorts,
	filters *poolPayoutsProto.MinedBlocksFilters,
	limit, offset uint32,
) (*poolPayoutsProto.MinedBlocksList, error) {
	client := poolPayoutsProto.NewPoolPayoutsServiceClient(blockchain.GetConnection())
	blocksList, err := client.GetBlocks(ctx, &poolPayoutsProto.GetBlocksRequest{
		Filters: filters,
		Sorts:   sorts,
		Pagination: &paginationProto.PaginationRequest{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) blocks: %w", blockchain.GetInfo().Coin, err)
	}

	return blocksList, nil
}

func (s *BlockchainService) GetSoloBlocks(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.MinedSoloBlocksSorts,
	filters *poolPayoutsProto.MinedSoloBlocksFilters,
	limit, offset uint32,
) (*poolPayoutsProto.MinedSoloBlocksList, error) {
	client := s.getProtoClient(blockchain)
	soloBlocksList, err := client.GetSoloBlocks(ctx, &poolPayoutsProto.GetSoloBlocksRequest{
		Filters: filters,
		Sorts:   sorts,
		Pagination: &paginationProto.PaginationRequest{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		switch status.Code(err) {
		case codes.Unimplemented:
			return nil, nil
		default:
			return nil, fmt.Errorf("failed to get blockchain (coin: %s) solo blocks: %w", blockchain.GetInfo().Coin, err)
		}
	}

	return soloBlocksList, nil
}
