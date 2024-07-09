package constant

import (
	"net/http"

	"github.com/Halalins/backend/internal/common/customerror"
)

var (
	ErrInternalServerError = &customerror.CustomError{
		HTTPCode:  http.StatusInternalServerError,
		Message:   "Internal Server Error",
		ErrorCode: "INTERNAL_SERVER_ERROR",
	}
	ErrDataNotFound = &customerror.CustomError{
		HTTPCode:  http.StatusNotFound,
		Message:   "Data Not Found",
		ErrorCode: "DATA_NOT_FOUND",
	}
	ErrUnauthorized = &customerror.CustomError{
		HTTPCode:  http.StatusUnauthorized,
		Message:   "Unauthorized",
		ErrorCode: "UNAUTHORIZED",
	}
	ErrFileTypeIsNotSupported = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "File type is not supported",
		ErrorCode: "FILE_TYPE_IS_NOT_SUPPORTED",
	}
	ErrActionNotAllowed = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Action not allowed, please check your payload request",
		ErrorCode: "ACTION_NOT_ALLOWED_PLEASE_CHECK_YOUR_PAYLOAD_REQUEST",
	}
	ErrDuplicateBoilerplate = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "boilerplate be more than 1 option",
		ErrorCode: "INVALID_REQUEST",
	}
	ErrInvalidRequest = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Invalid request",
		ErrorCode: "INVALID_REQUEST",
	}
	ErrRoleNameAlreadyExist = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Role name already exist",
		ErrorCode: "ROLE_NAME_ALREADY_EXIST",
	}
	ErrCanNotAssignSuperAdminRole = &customerror.CustomError{
		HTTPCode:  http.StatusNotFound,
		Message:   "Can Not Assign Super Admin Role",
		ErrorCode: "CAN_NOT_ASSIGN_SUPER_ADMIN_ROLE",
	}
	ErrBindingError = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Binding Validation Error",
		ErrorCode: "BINDING_VALIDATION_ERROR",
	}
	ErrInvalidUUID = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Invalid UUID",
		ErrorCode: "INVALID_UUID",
	}
	ErrTokenNotMatchOrInvalid = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Token Not Match or Invalid",
		ErrorCode: "TOKEN_NOT_MATCH_OR_INVALID",
	}
	ErrRoleNotFound = &customerror.CustomError{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Role not found",
		ErrorCode: "ROLE_NOT_FOUND",
	}
	ErrAccountAlreadyExist = &customerror.CustomError{
		HTTPCode:  http.StatusConflict,
		Message:   "Account Already Exist",
		ErrorCode: "ACCOUNT_ALREADY_EXIST",
	}
	ErrAccessPermissionDenied = &customerror.CustomError{
		HTTPCode:  http.StatusForbidden,
		Message:   "Access Permission Denied",
		ErrorCode: "ACCESS_PERMISSION_DENIED",
	}
	ErrMediaNotFound = &customerror.CustomError{
		HTTPCode:  http.StatusNotFound,
		Message:   "Media Not Found",
		ErrorCode: "MEDIA_NOT_FOUND",
	}
	ErrBankNotFound = &customerror.CustomError{
		HTTPCode:  http.StatusNotFound,
		Message:   "Bank Not Found",
		ErrorCode: "BANK_NOT_FOUND",
	}
	ErrToManyRequest = &customerror.CustomError{
		HTTPCode:  http.StatusTooManyRequests,
		Message:   "To many request",
		ErrorCode: "TO_MANY_REQUEST",
	}
	ErrPinInvalid = &customerror.CustomError{
		HTTPCode:  http.StatusUnauthorized,
		Message:   "PIN is invalid",
		ErrorCode: "PIN_IS_INVALID",
	}
	ErrOTPInvalid = &customerror.CustomError{
		HTTPCode:  http.StatusUnauthorized,
		Message:   "OTP is invalid",
		ErrorCode: "OTP_IS_INVALID",
	}
)

var (
	MapConstraintError = map[string]*customerror.CustomError{
		"boilerplate_key":                              ErrDuplicateBoilerplate,
		"admins_role_id_fkey":                          ErrRoleNotFound,
		"unique constraint \"admins_email_key\"":       ErrAccountAlreadyExist,
		"unique constraint \"admins_username_key\"":    ErrAccountAlreadyExist,
		"foreign key constraint \"admins_photo_fkey\"": ErrMediaNotFound,
		"agents_bank_id_fkey":                          ErrBankNotFound,
		"agents_npwp_document_fkey":                    ErrMediaNotFound,
		"agents_ktp_document_fkey":                     ErrMediaNotFound,
		"agents_email_key":                             ErrAccountAlreadyExist,
		"agents_username_key":                          ErrAccountAlreadyExist,
	}
)
