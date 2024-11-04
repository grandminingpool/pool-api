package blocksErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	GetBlocksError                   serverErrors.ServerErrorCode = "get_blocks_error"
	GetSoloBlocksError               serverErrors.ServerErrorCode = "get_solo_blocks_error"
	GetSoloBlocksNotImplementedError serverErrors.ServerErrorCode = "get_solo_blocks_not_implemented"
)

func CreateGetBlocksError(err error) *apiModels.GetBlockchainBlocksInternalServerError {
	return &apiModels.GetBlockchainBlocksInternalServerError{
		Code:    string(GetBlocksError),
		Message: err.Error(),
	}
}

func CreateGetSoloBlocksError(err error) *apiModels.GetBlockchainSoloBlocksInternalServerError {
	return &apiModels.GetBlockchainSoloBlocksInternalServerError{
		Code:    string(GetSoloBlocksError),
		Message: err.Error(),
	}
}

func CreateGetSoloBlocksNotImplementedError() *apiModels.GetBlockchainSoloBlocksMethodNotAllowed {
	return &apiModels.GetBlockchainSoloBlocksMethodNotAllowed{
		Code:    string(GetSoloBlocksNotImplementedError),
		Message: "Getting solo blocks list is not implemented",
	}
}
