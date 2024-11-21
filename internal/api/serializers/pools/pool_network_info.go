package poolsSerializers

import (
	"context"
	"math/big"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolNetworkInfoSerialzier struct{}

func (s *PoolNetworkInfoSerialzier) Serialize(ctx context.Context, networkInfo *poolProto.NetworkInfo) *apiModels.PoolNetworkInfo {
	return &apiModels.PoolNetworkInfo{
		TopBlockHash: networkInfo.TopBlockHash,
		Difficulty:   new(big.Int).SetBytes(networkInfo.Difficulty).String(),
		BlockReward:  networkInfo.BlockReward,
	}
}
