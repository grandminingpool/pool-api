package blocksQuery

import (
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParseSoloBlocksFiltersInQueryParams(params *apiModels.GetBlockchainSoloBlocksParams) *poolPayoutsProto.MinedSoloBlocksFilters {
	filters := &poolPayoutsProto.MinedSoloBlocksFilters{}

	if params.Miner.IsSet() {
		filters.Miner = &params.Miner.Value
	}

	if params.MinerHashrate.IsSet() {
		filters.MinerHashrate = poolAPIQueryUtils.ParseBigIntRangeFilter(params.MinerHashrate.Value)
	}

	if params.BlockHash.IsSet() {
		filters.BlockHash = &params.BlockHash.Value
	}

	if params.Reward.IsSet() {
		filters.Reward = poolAPIQueryUtils.ParseUInt64RangeFilter(params.Reward.Value)
	}

	if params.TxHash.IsSet() {
		filters.TxHash = &params.TxHash.Value
	}

	if params.ShareDifficulty.IsSet() {
		filters.ShareDifficulty = poolAPIQueryUtils.ParseUInt64RangeFilter(params.ShareDifficulty.Value)
	}

	if params.MinedAt.IsSet() {
		filters.MinedAt = poolAPIQueryUtils.ParseDateTimeRangeFilter(params.MinedAt.Value)
	}

	return filters
}
