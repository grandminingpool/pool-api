package minersSerializer

import (
	"context"
	"math/big"
	"time"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinerWorkerSerializer struct{}

func (s *MinerWorkerSerializer) Serialize(ctx context.Context, minerWorker *poolMinersProto.MinerWorker) *apiModels.MinerWorker {
	return &apiModels.MinerWorker{
		Worker:      minerWorker.Worker,
		Region:      minerWorker.Region,
		Agent:       minerWorker.Agent,
		Solo:        minerWorker.Solo,
		Hashrate:    new(big.Int).SetBytes(minerWorker.Hashrate).String(),
		ConnectedAt: minerWorker.ConnectedAt.AsTime().Format(time.RFC3339),
	}
}
