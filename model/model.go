package model

type BaseError struct {
	Code     int         `json:"code"`
	CodeText string      `json:"codeText"`
	Message  string      `json:"message"`
	ErrorMsg interface{} `json:"errorMsg"`
}

func (b BaseError) Error() string {
	return b.Message
}

type CustomError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

type EntError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

type OauthError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}

type OtelError struct {
	ErrorId  string `json:"errorId"`
	Domain   string `json:"domain"`
	ErrorMsg error  `json:"errorMsg"`
}
