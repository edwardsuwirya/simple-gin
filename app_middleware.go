package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	c.Next()
}
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()

		errResp := ErrorMessage{}
		err := json.Unmarshal([]byte(e), &errResp)
		if err != nil {
			errResp.HttpCode = 500
			errResp.ErrorDescription = ErrorDescription{
				Code:        "06",
				Description: "Convert json failed",
			}
		}
		NewJsonResponse(c).SendError(errResp)
	}
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")

		if token == "" {
			c.AbortWithStatusJSON(401, "Unauthorized")
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, "Unauthorized")
			return
		}

		c.Next()
	}
}
