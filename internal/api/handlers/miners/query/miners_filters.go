package minersQuery

import (
	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParseMinersFiltersInQueryParams(params *apiModels.GetBlockchainMinersParams) *poolMinersProto.MinersFilters {
	filters := &poolMinersProto.MinersFilters{}

	if params.Address.IsSet() {
		filters.Address = &params.Address.Value
	}

	if params.Hashrate.IsSet() {
		filters.Hashrate = poolAPIQueryUtils.ParseBigIntRangeFilter(params.Hashrate.Value)
	}

	if params.WorkersCount.IsSet() {
		filters.WorkersCount = poolAPIQueryUtils.ParseUInt32Filter(params.WorkersCount.Value)
	}

	if params.BlocksCount.IsSet() {
		filters.BlocksCount = poolAPIQueryUtils.ParseUInt32Filter(params.BlocksCount.Value)
	}

	if params.SoloBlocksCount.IsSet() {
		filters.SoloBlocksCount = poolAPIQueryUtils.ParseUInt32Filter(params.SoloBlocksCount.Value)
	}

	if params.JoinedAt.IsSet() {
		filters.JoinedAt = poolAPIQueryUtils.ParseDateTimeRangeFilter(params.JoinedAt.Value)
	}

	return filters
}
