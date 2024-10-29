package payoutsQuery

import (
	poolPayoutsProto "github.com/grandminingpool/pool-api-proto/generated/pool_payouts"
	sortsProto "github.com/grandminingpool/pool-api-proto/generated/utils/sorts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolAPIQueryUtils "github.com/grandminingpool/pool-api/internal/common/utils/pool_api/query"
)

func ParsePayoutsSortsInQuery(querySorts *apiModels.OptString) *poolPayoutsProto.PayoutsSorts {
	sorts := &poolPayoutsProto.PayoutsSorts{
		PaidAt: &sortsProto.SortOrder{
			Direction: sortsProto.SortDirection_DESC,
		},
	}

	if querySorts.IsSet() {
		querySortsMap := poolAPIQueryUtils.ParseSortsItems(querySorts.Value)

		if miner, ok := querySortsMap["miner"]; ok {
			sorts.Miner = miner
		}

		if txHash, ok := querySortsMap["tx_hash"]; ok {
			sorts.TxHash = txHash
		}

		if amount, ok := querySortsMap["amount"]; ok {
			sorts.Amount = amount
		}

		if paidAt, ok := querySortsMap["paid_at"]; ok {
			sorts.PaidAt = paidAt
		}
	}

	return sorts
}
