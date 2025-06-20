package errorx

import (
	"lib/model"
	"net/http"

	"lib/ent/entgen"
)

type EntError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

func (err EntError) Error() string {
	return err.ErrorMsg.Error()
}

func HandleEntError(err *EntError) model.BaseError {
	var baseErr model.BaseError

	switch e := err.ErrorMsg.(type) {
	case *entgen.ConstraintError:
		baseErr.Code = http.StatusBadRequest
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
		baseErr.Message = "Constraint violation: " + e.Error()
		baseErr.ErrorMsg = e

	case *entgen.NotFoundError:
		baseErr.Code = http.StatusNotFound
		baseErr.CodeText = http.StatusText(http.StatusNotFound)
		baseErr.Message = "Resource not found: " + e.Error()
		baseErr.ErrorMsg = e

	case *entgen.ValidationError:
		baseErr.Code = http.StatusBadRequest
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
		baseErr.Message = "Validation failed: " + e.Error()
		baseErr.ErrorMsg = e

	default:
		baseErr = SetDefaultError(e)
	}

	return baseErr
}
