package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	ctx *gin.Context
}

func (j *JsonResponse) SendData(message ResponseMessage) {
	j.ctx.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(errMessage ErrorMessage) {
	j.ctx.AbortWithStatusJSON(errMessage.HttpCode, errMessage)
}

func NewJsonResponse(ctx *gin.Context) AppHttpResponse {
	return &JsonResponse{ctx: ctx}
}
