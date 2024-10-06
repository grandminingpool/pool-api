package blockchainsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainSerializer struct{}

func (s *BlockchainSerializer) Serialize(ctx context.Context, blockchainInfo *blockchains.BlockchainInfo) *apiModels.Blockchain {
	return &apiModels.Blockchain{
		Coin:       blockchainInfo.Coin,
		Name:       blockchainInfo.Name,
		Ticker:     blockchainInfo.Ticker,
		AtomicUnit: int(blockchainInfo.AtomicUnit),
	}
}
