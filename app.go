package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	productRoute := r.Group("/product")
	productRoute.GET("", gettingWithQueryParam())
	productRoute.GET("/:id", gettingWithPathVariable())
	productRoute.POST("", posting())

	err := r.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}

func gettingWithQueryParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNo := c.Query("page")
		itemPerPage := c.Query("itempage")
		c.JSON(200, gin.H{
			"pageNo":      pageNo,
			"itemPerPage": itemPerPage,
		})
	}
}
func gettingWithPathVariable() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("id")
		c.JSON(200, gin.H{
			"message": productId,
		})
	}
}
func posting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productReq ProductRequest
		if err := c.ShouldBindJSON(&productReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message": productReq,
		})
	}
}
