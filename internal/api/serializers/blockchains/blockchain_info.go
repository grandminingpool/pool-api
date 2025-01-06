package blockchainsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainInfoSerializer struct{}

func (s *BlockchainInfoSerializer) Serialize(ctx context.Context, blockchainInfo blockchains.BlockchainInfo) apiModels.BlockchainInfo {
	return apiModels.BlockchainInfo{
		Blockchain: blockchainInfo.Blockchain,
		Name:       blockchainInfo.Name,
		Ticker:     blockchainInfo.Ticker,
		AtomicUnit: blockchainInfo.AtomicUnit,
	}
}
