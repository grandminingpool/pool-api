package blocksQuery

import (
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParseBlocksFiltersInQueryParams(params *apiModels.GetBlockchainBlocksParams) *poolPayoutsProto.MinedBlocksFilters {
	filters := &poolPayoutsProto.MinedBlocksFilters{}

	if params.Miner.IsSet() {
		filters.Miner = &params.Miner.Value
	}

	if params.MinerHashrate.IsSet() {
		filters.MinerHashrate = poolAPIQueryUtils.ParseBigIntRangeFilter(params.MinerHashrate.Value)
	}

	if params.BlockHash.IsSet() {
		filters.BlockHash = &params.BlockHash.Value
	}

	if params.RoundMinersCount.IsSet() {
		filters.RoundMinersCount = poolAPIQueryUtils.ParseUInt32Filter(params.RoundMinersCount.Value)
	}

	if params.MinedAt.IsSet() {
		filters.MinedAt = poolAPIQueryUtils.ParseDateTimeRangeFilter(params.MinedAt.Value)
	}

	return filters
}
