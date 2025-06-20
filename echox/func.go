package echox

import (
	"lib/errorx"
	"lib/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HttpErrorHandler(err error, c echo.Context) {
	var errResponse model.BaseError

	switch t := err.(type) {
	case *errorx.CustomError:
		errResponse = errorx.HandleCustomError(t)

	case *errorx.OauthError:
		errResponse = errorx.HandleOauthError(t)

	case *errorx.EntError:
		errResponse = errorx.HandleEntError(t)

	default:
		log.Printf("Unknown error: %v", err)
		errResponse = setDefaultError(err)
	}

	if responseErr := c.JSON(errResponse.Code, getErrorResponse(errResponse)); responseErr != nil {
		c.Logger().Error(responseErr)
	}
}

type BaseStructure struct {
	ApiVersion string      `json:"apiVersion,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func setDefaultError(err error) model.BaseError {
	return model.BaseError{
		Code:     http.StatusInternalServerError,
		CodeText: http.StatusText(http.StatusInternalServerError),
		Message:  "An internal server error occurred",
		ErrorMsg: err,
	}
}

func getErrorResponse(errResponse model.BaseError) BaseStructure {
	return BaseStructure{
		Error: errResponse,
	}
}
