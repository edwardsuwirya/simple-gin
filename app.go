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
	r.Use(DummyMiddleware)
	r.Use(TokenAuthMiddleware())
	productRoute := r.Group("/product")
	productRoute.Use(DummyMiddleware)
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
		NewJsonResponse(c).SendData(NewResponseMessage("00", "Product ID", productId))
	}
}
func posting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productReq ProductRequest
		if err := c.ShouldBindJSON(&productReq); err != nil {
			NewJsonResponse(c).SendError(NewErrorMessage(http.StatusBadRequest, "01", err.Error()))
			return
		}
		NewJsonResponse(c).SendData(NewResponseMessage("00", "Create Product", productReq))
	}
}
