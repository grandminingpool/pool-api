package payoutsErrors

import serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"

const (
	GetPayoutsError           serverErrors.ServerErrorCode = "get_payouts_error"
	GetMinerBalanceError      serverErrors.ServerErrorCode = "get_miner_balance_error"
	MinerBalanceNotFoundError serverErrors.ServerErrorCode = "miner_balance_not_found"
)
