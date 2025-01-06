package blocksQuery

import (
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	sortsProto "github.com/grandminingpool/pool-api-proto/generated/utils/sorts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParseSoloBlocksSortsInQuery(querySorts *apiModels.OptString) *poolPayoutsProto.MinedSoloBlocksSorts {
	sorts := &poolPayoutsProto.MinedSoloBlocksSorts{
		MinedAt: &sortsProto.SortOrder{
			Direction: sortsProto.SortDirection_DESC,
		},
	}

	if querySorts.IsSet() {
		querySortsMap := poolAPIQueryUtils.ParseSortsItems(querySorts.Value)

		if miner, ok := querySortsMap["miner"]; ok {
			sorts.Miner = miner
		}

		if minerHashrate, ok := querySortsMap["miner_hashrate"]; ok {
			sorts.MinerHashrate = minerHashrate
		}

		if blockHash, ok := querySortsMap["block_hash"]; ok {
			sorts.BlockHash = blockHash
		}

		if reward, ok := querySortsMap["reward"]; ok {
			sorts.Reward = reward
		}

		if txHash, ok := querySortsMap["tx_hash"]; ok {
			sorts.TxHash = txHash
		}

		if minedAt, ok := querySortsMap["mined_at"]; ok {
			sorts.MinedAt = minedAt
		}
	}

	return sorts
}
