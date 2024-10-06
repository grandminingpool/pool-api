package poolsErrors

import serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"

const (
	GetPoolDataError  serverErrors.ServerErrorCode = "get_pool_data_error"
	GetPoolInfoError  serverErrors.ServerErrorCode = "get_pool_info_error"
	GetPoolStatsError serverErrors.ServerErrorCode = "get_pool_stats_error"
)
