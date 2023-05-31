package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	fmt.Printf("authHeader")

	if authHeader == "" {
		c.Next()
	}

	fmt.Printf(authHeader)
}
