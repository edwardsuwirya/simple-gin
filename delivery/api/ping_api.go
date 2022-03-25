package api

import "github.com/gin-gonic/gin"

type PingApi struct {
}

func (p *PingApi) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func NewPingApi(pingRoute *gin.RouterGroup) {
	pingApi := PingApi{}
	pingRoute.GET("", pingApi.Ping())
}
