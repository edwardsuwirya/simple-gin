package main

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
