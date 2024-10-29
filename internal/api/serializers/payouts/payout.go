package payoutsSerializer

import (
	"context"
	"time"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PayoutSerializer struct{}

func (s *PayoutSerializer) Serialize(ctx context.Context, payout *poolPayoutsProto.Payout) *apiModels.Payout {
	return &apiModels.Payout{
		Miner:  payout.Miner,
		TxHash: payout.TxHash,
		Amount: payout.Amount,
		PaidAt: payout.PaidAt.AsTime().Format(time.RFC3339),
	}
}
