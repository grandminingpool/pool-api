package minersQuery

import (
	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	sortsProto "github.com/grandminingpool/pool-api-proto/generated/utils/sorts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParseMinersSortsInQuery(querySorts *apiModels.OptString) *poolMinersProto.MinersSorts {
	sorts := &poolMinersProto.MinersSorts{
		JoinedAt: &sortsProto.SortOrder{
			Direction: sortsProto.SortDirection_DESC,
		},
	}

	if querySorts.IsSet() {
		querySortsMap := poolAPIQueryUtils.ParseSortsItems(querySorts.Value)

		if address, ok := querySortsMap["address"]; ok {
			sorts.Address = address
		}

		if hashrate, ok := querySortsMap["hashrate"]; ok {
			sorts.Hashrate = hashrate
		}

		if workersCount, ok := querySortsMap["workers_count"]; ok {
			sorts.WorkersCount = workersCount
		}

		if blocksCount, ok := querySortsMap["blocks_count"]; ok {
			sorts.BlocksCount = blocksCount
		}

		if soloBlocksCount, ok := querySortsMap["solo_blocks_count"]; ok {
			sorts.SoloBlocksCount = soloBlocksCount
		}

		if joinedAt, ok := querySortsMap["joined_at"]; ok {
			sorts.JoinedAt = joinedAt
		}
	}

	return sorts
}
