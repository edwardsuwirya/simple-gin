package api

import (
	"enigmacamp.com/simplegin/delivery/apprequest"
	"enigmacamp.com/simplegin/delivery/commonresp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductApi struct {
}

func (p *ProductApi) gettingWithQueryParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNo := c.Query("page")
		itemPerPage := c.Query("itempage")
		fmt.Println("Exec")
		c.JSON(200, gin.H{
			"pageNo":      pageNo,
			"itemPerPage": itemPerPage,
		})
	}
}

func (p *ProductApi) gettingWithPathVariable() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("id")
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", "Product ID", productId))
	}
}

func (p *ProductApi) posting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productReq apprequest.ProductRequest
		if err := c.ShouldBindJSON(&productReq); err != nil {
			c.Error(fmt.Errorf("%s", commonresp.NewErrorMessage(http.StatusBadRequest, "01", err.Error()).ToJson()))
			return
		}
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", "Create Product", productReq))
	}
}

func NewProductApi(productRoute *gin.RouterGroup) {
	productApi := ProductApi{}
	productRoute.GET("", productApi.gettingWithQueryParam())
	productRoute.GET("/:id", productApi.gettingWithPathVariable())
	productRoute.POST("", productApi.posting())
}
