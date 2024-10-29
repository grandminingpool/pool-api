package blocksSerializers

import (
	"context"
	"math/big"
	"time"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinedSoloBlockSerializer struct{}

func (s *MinedSoloBlockSerializer) Serialize(ctx context.Context, minedSoloBlock *poolPayoutsProto.MinedSoloBlock) *apiModels.MinedSoloBlock {
	return &apiModels.MinedSoloBlock{
		Miner:           minedSoloBlock.Miner,
		MinerHashrate:   new(big.Int).SetBytes(minedSoloBlock.MinerHashrate).String(),
		BlockHash:       minedSoloBlock.BlockHash,
		Reward:          minedSoloBlock.Reward,
		TxHash:          minedSoloBlock.TxHash,
		ShareDifficulty: minedSoloBlock.ShareDifficulty,
		MinedAt:         minedSoloBlock.MinedAt.AsTime().Format(time.RFC3339),
	}
}
