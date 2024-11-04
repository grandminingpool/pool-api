package blocksHandlers

import (
	"context"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blocksErrors "github.com/grandminingpool/pool-api/internal/api/handlers/blocks/errors"
	blocksServices "github.com/grandminingpool/pool-api/internal/api/services/blocks"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	blockchainService        *blocksServices.BlockchainService
	minedBlockSerializer     serializers.BaseSerializer[*poolPayoutsProto.MinedBlock, *apiModels.MinedBlock]
	minedSoloBlockSerializer serializers.BaseSerializer[*poolPayoutsProto.MinedSoloBlock, *apiModels.MinedSoloBlock]
}

func (h *BlockchainHandler) GetBlocks(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.MinedBlocksSorts,
	filters *poolPayoutsProto.MinedBlocksFilters,
	limit, offset uint32,
) apiModels.GetBlockchainBlocksRes {
	blocksList, err := h.blockchainService.GetBlocks(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return blocksErrors.CreateGetBlocksError(err)
	}

	blocksResponse := make([]apiModels.MinedBlock, 0, len(blocksList.Blocks))
	for _, b := range blocksList.Blocks {
		blocksResponse = append(blocksResponse, *h.minedBlockSerializer.Serialize(ctx, b))
	}

	return &apiModels.MinedBlocksList{
		Blocks: blocksResponse,
		Limit:  blocksList.Pagination.Limit,
		Offset: blocksList.Pagination.Offset,
		Total:  blocksList.Pagination.Total,
	}
}

func (h *BlockchainHandler) GetSoloBlocks(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.MinedSoloBlocksSorts,
	filters *poolPayoutsProto.MinedSoloBlocksFilters,
	limit, offset uint32,
) apiModels.GetBlockchainSoloBlocksRes {
	soloBlocksList, err := h.blockchainService.GetSoloBlocks(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return blocksErrors.CreateGetSoloBlocksError(err)
	} else if soloBlocksList == nil {
		return blocksErrors.CreateGetSoloBlocksNotImplementedError()
	}

	soloBlocksResponse := make([]apiModels.MinedSoloBlock, 0, len(soloBlocksList.Blocks.Blocks))
	for _, b := range soloBlocksList.Blocks.Blocks {
		soloBlocksResponse = append(soloBlocksResponse, *h.minedSoloBlockSerializer.Serialize(ctx, b))
	}

	return &apiModels.MinedSoloBlocksList{
		Blocks: soloBlocksResponse,
		Limit:  soloBlocksList.Pagination.Limit,
		Offset: soloBlocksList.Pagination.Offset,
		Total:  soloBlocksList.Pagination.Total,
	}
}

func NewBlockchainHandler(
	blockchainService *blocksServices.BlockchainService,
	minedBlockSerializer serializers.BaseSerializer[*poolPayoutsProto.MinedBlock, *apiModels.MinedBlock],
	minedSoloBlockSerializer serializers.BaseSerializer[*poolPayoutsProto.MinedSoloBlock, *apiModels.MinedSoloBlock],
) *BlockchainHandler {
	return &BlockchainHandler{
		blockchainService:        blockchainService,
		minedBlockSerializer:     minedBlockSerializer,
		minedSoloBlockSerializer: minedSoloBlockSerializer,
	}
}
