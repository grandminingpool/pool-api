package payoutsHandlers

import (
	"context"
	"fmt"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	payoutsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/payouts/errors"
	payoutsServices "github.com/grandminingpool/pool-api/internal/api/services/payouts"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
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
) (*apiModels.PayoutsList, error) {
	payoutsList, err := h.blockchainService.GetPayouts(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(payoutsErrors.GetPayoutsError, err)
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
	}, nil
}

func (h *BlockchainHandler) GetMinerBalance(ctx context.Context, blockchain *blockchains.Blockchain, miner string) (*apiModels.MinerBalance, error) {
	minerBalance, err := h.blockchainService.GetMinerBalance(ctx, blockchain, miner)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(payoutsErrors.GetMinerBalanceError, err)
	} else if minerBalance == nil {
		return nil, serverErrors.CreateNotFoundError(payoutsErrors.MinerBalanceNotFoundError, fmt.Errorf("miner '%s' balance not found", miner))
	}

	return &apiModels.MinerBalance{
		Balance: *minerBalance,
	}, nil
}
