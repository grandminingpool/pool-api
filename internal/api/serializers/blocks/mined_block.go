package blocksSerializers

import (
	"context"
	"math/big"
	"time"

	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinedBlockSerializer struct{}

func (s *MinedBlockSerializer) Serialize(ctx context.Context, minedBlock *poolPayoutsProto.MinedBlock) *apiModels.MinedBlock {
	return &apiModels.MinedBlock{
		Miner:            minedBlock.Miner,
		MinerHashrate:    new(big.Int).SetBytes(minedBlock.MinerHashrate).String(),
		BlockHash:        minedBlock.BlockHash,
		RoundMinersCount: minedBlock.RoundMinersCount,
		MinedAt:          minedBlock.MinedAt.AsTime().Format(time.RFC3339),
	}
}
