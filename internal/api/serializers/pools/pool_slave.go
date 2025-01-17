package poolsSerializers

import (
	"context"
	"time"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolSlaveSerialzier struct{}

func (s *PoolSlaveSerialzier) Serialize(ctx context.Context, poolSlave *poolProto.PoolSlave) apiModels.PoolSlave {
	return apiModels.PoolSlave{
		Region:      poolSlave.Region,
		Host:        poolSlave.Host,
		TCPPort:     poolSlave.TcpPort,
		SslPort:     poolSlave.SslPort,
		ConnectedAt: poolSlave.ConnectedAt.AsTime().Format(time.RFC3339),
	}
}
