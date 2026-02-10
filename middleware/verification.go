package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
		return
	}
	if !strings.HasPrefix(token, "Bearer ") {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	if parsedToken.Valid {
		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
}
