package errorx

import (
	"lib/model"
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
)

type OauthError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

func (err *OauthError) Error() string {
	return err.ErrorMsg.Error()
}

func HandleOauthError(err *OauthError) model.BaseError {
	var baseErr model.BaseError

	switch err.ErrorMsg {
	case errors.ErrInvalidAccessToken:
		baseErr.Code = http.StatusUnauthorized // 401
		baseErr.CodeText = http.StatusText(http.StatusUnauthorized)
		baseErr.Message = "Invalid access token"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrExpiredAccessToken:
		baseErr.Code = http.StatusUnauthorized // 401
		baseErr.CodeText = http.StatusText(http.StatusUnauthorized)
		baseErr.Message = "Access token has expired"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrAccessDenied:
		baseErr.Code = http.StatusForbidden // 403
		baseErr.CodeText = http.StatusText(http.StatusForbidden)
		baseErr.Message = "Access denied: insufficient permissions"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrInvalidGrant:
		baseErr.Code = http.StatusBadRequest // 400
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
		baseErr.Message = "Invalid authorization grant"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrUnsupportedGrantType:
		baseErr.Code = http.StatusBadRequest // 400
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
		baseErr.Message = "Unsupported grant type"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrInvalidRefreshToken:
		baseErr.Code = http.StatusUnauthorized // 401
		baseErr.CodeText = http.StatusText(http.StatusUnauthorized)
		baseErr.Message = "Invalid or expired refresh token"
		baseErr.ErrorMsg = err.ErrorMsg

	case errors.ErrInvalidScope:
		baseErr.Code = http.StatusBadRequest // 400
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
		baseErr.Message = "Invalid scope in request"
		baseErr.ErrorMsg = err.ErrorMsg

	default:
		log.Println("Handling OAuth default error case")
		baseErr = SetOauthDefaultError(err)
	}

	return baseErr
}

func SetOauthDefaultError(err *OauthError) model.BaseError {
	return model.BaseError{
		Code:     http.StatusInternalServerError,
		CodeText: http.StatusText(http.StatusInternalServerError),
		Message:  "An unexpected error occurred during the OAuth process.",
		ErrorMsg: err.Error(),
	}
}
