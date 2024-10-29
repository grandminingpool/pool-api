package minersSerializer

import (
	"context"
	"math/big"
	"time"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinerSerializer struct{}

func (s *MinerSerializer) Serialize(ctx context.Context, miner *poolMinersProto.Miner) *apiModels.Miner {
	return &apiModels.Miner{
		Address:         miner.Address,
		Hashrate:        new(big.Int).SetBytes(miner.Hashrate).String(),
		WorkersCount:    miner.WorkersCount,
		BlocksCount:     miner.BlocksCount,
		SoloBlocksCount: miner.SoloBlocksCount,
		JoinedAt:        miner.JoinedAt.AsTime().Format(time.RFC3339),
		LastActivity:    miner.LastActivity.AsTime().Format(time.RFC3339),
	}
}
