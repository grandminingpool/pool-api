package payoutsQuery

import (
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParsePayoutsFiltersInQueryParams(params *apiModels.GetBlockchainPayoutsParams) *poolPayoutsProto.PayoutsFilters {
	filters := &poolPayoutsProto.PayoutsFilters{}

	if params.Miner.IsSet() {
		filters.Miner = &params.Miner.Value
	}

	if params.TxHash.IsSet() {
		filters.TxHash = &params.TxHash.Value
	}

	if params.Amount.IsSet() {
		filters.Amount = poolAPIQueryUtils.ParseUInt64RangeFilter(params.Amount.Value)
	}

	if params.PaidAt.IsSet() {
		filters.PaidAt = poolAPIQueryUtils.ParseDateTimeRangeFilter(params.PaidAt.Value)
	}

	return filters
}
