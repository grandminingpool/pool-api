package minersErrors

import serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"

const (
	GetMinersError       serverErrors.ServerErrorCode = "get_miners_error"
	GetMinerError        serverErrors.ServerErrorCode = "get_miner_error"
	MinerNotFoundError   serverErrors.ServerErrorCode = "miner_not_found"
	GetMinerWorkersError serverErrors.ServerErrorCode = "get_miner_workers_error"
)
