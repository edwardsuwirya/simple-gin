package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	c.Next()
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
