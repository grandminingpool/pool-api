package payoutsHandlers

import (
	"context"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	payoutsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/payouts/errors"
	payoutsServices "github.com/grandminingpool/pool-api/internal/api/services/payouts"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	blockchainService *payoutsServices.BlockchainService
	payoutSerializer  serializers.BaseSerializer[*poolPayoutsProto.Payout, *apiModels.Payout]
}

func (h *BlockchainHandler) GetPayouts(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolPayoutsProto.PayoutsSorts,
	filters *poolPayoutsProto.PayoutsFilters,
	limit, offset uint32,
) apiModels.GetBlockchainPayoutsRes {
	payoutsList, err := h.blockchainService.GetPayouts(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return payoutsErrors.CreateGetPayoutsError(err)
	}

	payoutsResponse := make([]apiModels.Payout, 0, len(payoutsList.Payouts.Payouts))
	for _, p := range payoutsList.Payouts.Payouts {
		payoutsResponse = append(payoutsResponse, *h.payoutSerializer.Serialize(ctx, p))
	}

	return &apiModels.PayoutsList{
		Payouts: payoutsResponse,
		Limit:   payoutsList.Pagination.Limit,
		Offset:  payoutsList.Pagination.Offset,
		Total:   payoutsList.Pagination.Total,
	}
}

func (h *BlockchainHandler) GetMinerBalance(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	miner string,
) apiModels.GetBlockchainMinerBalanceRes {
	minerBalance, err := h.blockchainService.GetMinerBalance(ctx, blockchain, miner)
	if err != nil {
		return payoutsErrors.CreateGetMinerBalanceError(err)
	} else if minerBalance == nil {
		return payoutsErrors.CreateMinerBalanceNotFoundError(miner)
	}

	return &apiModels.MinerBalance{
		Balance: *minerBalance,
	}
}

func NewBlockchainHandler(
	blockchainService *payoutsServices.BlockchainService,
	payoutSerializer serializers.BaseSerializer[*poolPayoutsProto.Payout, *apiModels.Payout],
) *BlockchainHandler {
	return &BlockchainHandler{
		blockchainService: blockchainService,
		payoutSerializer:  payoutSerializer,
	}
}
