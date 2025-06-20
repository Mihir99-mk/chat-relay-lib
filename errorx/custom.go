package errorx

import (
	"log"
	"net/http"

	"github.com/Mihir99-mk/chat-relay-lib/model"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

func (err CustomError) Error() string {
	return err.ErrorMsg.Error()
}

func HandleCustomError(err *CustomError) model.BaseError {
	var baseErr model.BaseError

	switch e := err.ErrorMsg.(type) {
	case *echo.HTTPError:
		baseErr.Code = e.Code
		baseErr.CodeText = http.StatusText(e.Code)
		baseErr.ErrorMsg = e

	default:
		log.Println("Handling default error case")

		baseErr = SetDefaultError(e)
	}

	return baseErr
}

func SetDefaultError(err error) model.BaseError {
	return model.BaseError{
		Code:     http.StatusInternalServerError,
		CodeText: http.StatusText(http.StatusInternalServerError),
		Message:  "An error occurred",
		ErrorMsg: err.Error(),
	}
}
