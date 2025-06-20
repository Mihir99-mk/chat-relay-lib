package errorx

import (
	"log"
	"net/http"

	"github.com/Mihir99-mk/chat-relay-lib/model"

	"github.com/gorilla/websocket"
)

type WsError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

func (err *WsError) Error() error {
	return err.ErrorMsg
}

func HandleWsError(err *WsError) model.BaseError {
	var baseErr model.BaseError

	switch err.ErrorMsg {
	case websocket.ErrBadHandshake:
		baseErr.Code = http.StatusBadGateway
		baseErr.CodeText = http.StatusText(http.StatusBadGateway)
		baseErr.Message = "WebSocket handshake failed with upstream server"
		baseErr.ErrorMsg = err.ErrorMsg

	case websocket.ErrReadLimit:
		baseErr.Code = http.StatusRequestEntityTooLarge
		baseErr.CodeText = http.StatusText(http.StatusRequestEntityTooLarge)
		baseErr.Message = "WebSocket read limit exceeded"
		baseErr.ErrorMsg = err.ErrorMsg

	case websocket.ErrCloseSent:
		baseErr.Code = http.StatusServiceUnavailable
		baseErr.CodeText = http.StatusText(http.StatusServiceUnavailable)
		baseErr.Message = "WebSocket connection already closed"
		baseErr.ErrorMsg = err.ErrorMsg

	default:
		log.Println("Handling Web socket default error case")
		baseErr = SetWsDefaultError(err)
	}

	return baseErr
}

func SetWsDefaultError(err *WsError) model.BaseError {
	return model.BaseError{
		Code:     http.StatusInternalServerError,
		CodeText: http.StatusText(http.StatusInternalServerError),
		Message:  "Unexpected WebSocket error.",
		ErrorMsg: err.Error(),
	}
}
