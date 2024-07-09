package response

import "github.com/Halalins/backend/internal/common/customerror"

type ErrorResp struct {
	Error customerror.CustomError `json:"error"`
}
