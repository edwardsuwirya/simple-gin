package main

import (
	"encoding/json"
)

type AppHttpResponse interface {
	SendData(message ResponseMessage)
	SendError(errMessage ErrorMessage)
}
type ResponseMessage struct {
	Status       string      `json:"status"`
	ResponseCode string      `json:"responseCode"`
	Description  string      `json:"message"`
	Data         interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode         int
	ErrorDescription ErrorDescription
}

func (e ErrorMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

type ErrorDescription struct {
	Code        string `json:"errorCode"`
	Description string `json:"message"`
}

func NewResponseMessage(respCode string, description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		"Success", respCode, description, data,
	}
}

func NewErrorMessage(httpCode int, errCode string, message string) ErrorMessage {
	em := ErrorMessage{
		HttpCode: httpCode,
		ErrorDescription: ErrorDescription{
			Code:        errCode,
			Description: message,
		},
	}
	return em
}
