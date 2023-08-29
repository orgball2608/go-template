package error

import (
	"net/http"
)

type ErrorCode struct {
	HTTPCode    int          `json:"-"`
	Type        ErrorType    `json:"type"`
	Code        Code         `json:"code,omitempty"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

// Authentication
var (
	ErrAuthenticationTokenInvalid ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeAuthentication,
		Code:     CodeAuthenticationTokenInvalid,
	}

	ErrAuthenticationTokenExpired ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeAuthentication,
		Code:     CodeAuthenticationTokenExpired,
	}
	ErrAuthenticationTokenSfdcExpired ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeAuthentication,
		Code:     CodeAuthenticationTokenSfdcExpired,
	}
)

// General
var (
	ErrGeneralSomethingWentWrong ErrorCode = ErrorCode{
		HTTPCode: http.StatusInternalServerError,
		Type:     ErrorTypeLogic,
		Code:     CodeGeneralSomethingWentWrong,
	}

	ErrGeneralForbidden ErrorCode = ErrorCode{
		HTTPCode: http.StatusForbidden,
		Type:     ErrorTypeLogic,
		Code:     CodeGeneralForbidden,
	}
	ErrGeneralInvalidIp ErrorCode = ErrorCode{
		HTTPCode: http.StatusForbidden,
		Type:     ErrorTypeLogic,
		Code:     CodeGeneralInvalidIp,
	}
)

// Logic
var (
	ErrLogicLoginFailed ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicLoginFailed,
	}

	ErrLogicDowLoadFileFailed ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicDownloadFileFailed,
	}

	ErrLogicLoginFailedManyTimes ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicLoginFailedManyTimes,
	}

	ErrLogicRecordAlreadyExists ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicRecordAlreadyExists,
	}

	ErrLogicRecordNotFound ErrorCode = ErrorCode{
		HTTPCode: http.StatusNotFound,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicRecordNotFound,
	}

	ErrLogicFormatRequestParam ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicFormatRequestParam,
	}

	ErrLogicUseCurrentPasswordToChangePassword ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicUseCurrentPasswordToChangePassword,
	}

	ErrLogicPasswordIncorrect ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicPasswordIncorrect,
	}

	ErrLogicTokenInvalid ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicTokenInvalid,
	}

	ErrLogicTokenExpired ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicTokenExpired,
	}

	ErrLogicRequestResetPasswordExisted ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicRequestResetPasswordExisted,
	}

	ErrLogicRefreshTokenInvalid ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicRefreshTokenInvalid,
	}

	ErrLogicGroupIdExisted ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeGroupExisted,
	}

	ErrLogicUploadFileFail ErrorCode = ErrorCode{
		HTTPCode: http.StatusInternalServerError,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicUploadFileFail,
	}

	ErrContractNotExist ErrorCode = ErrorCode{
		HTTPCode: http.StatusNotFound,
		Type:     ErrorTypeLogic,
		Code:     CodeContractNotExist,
	}

	ErrLogicWriteValuesAreNotConsistent ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicWriteValuesAreNotConsistent,
	}

	ErrLogicUserExisted ErrorCode = ErrorCode{
		HTTPCode: http.StatusBadRequest,
		Type:     ErrorTypeLogic,
		Code:     CodeUserExisted,
	}

	ErrCreateFaqBookmark ErrorCode = ErrorCode{
		HTTPCode: http.StatusForbidden,
		Type:     ErrorTypeLogic,
		Code:     CodeCreateFaqBookmark,
	}

	ErrLogicSyncFileFaqFailed ErrorCode = ErrorCode{
		HTTPCode: http.StatusUnauthorized,
		Type:     ErrorTypeLogic,
		Code:     CodeLogicDownloadFileFailed,
	}
)
