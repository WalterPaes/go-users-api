package middlewares

import (
	"net/http"
	"strings"

	"github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(jwtAuth *jwt.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenString = strings.Split(bearerToken, " ")[1]
		}

		isValid := jwtAuth.ValidateToken(tokenString)
		if !isValid {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
