package wrapperx

import (
	"log"
	"net/http"

	"github.com/Mihir99-mk/chat-relay-lib/errorx"
	"github.com/Mihir99-mk/chat-relay-lib/model"
	"github.com/google/uuid"
)

func WrapToBaseError(code int, err error, msg string) error {
	return &model.BaseError{
		Code:     code,
		CodeText: http.StatusText(code),
		ErrorMsg: err.Error(),
		Message:  msg,
	}
}

func WrapToCustomError(domain string, err error) error {
	log.Println("custom init")
	return &errorx.CustomError{
		ErrorId:  uuid.New().String(),
		Domain:   domain,
		ErrorMsg: err,
	}
}

func WrapToEntError(domain string, err error) error {
	log.Println("custom init")
	return &errorx.EntError{
		ErrorId:  uuid.New().String(),
		Domain:   domain,
		ErrorMsg: err,
	}
}

func WrapToOauthError(domain string, err error) error {
	return &errorx.OauthError{
		ErrorId:  uuid.New().String(),
		Domain:   domain,
		ErrorMsg: err,
	}
}
