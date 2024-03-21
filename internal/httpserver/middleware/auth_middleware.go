package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/jwtutils"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JWTUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JWTUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authStr := c.Request.Header.Get("Authorization")
		if authStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "please provide a token",
			})
			return
		}

		authStrs := strings.Split(authStr, " ")
		if len(authStrs) != 2 || authStrs[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "invalid token - malformed",
			})
			return
		}

		tokenString := authStrs[1]
		claims, err := m.jwtUtil.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("invalid token - %v", err.Error()),
			})
			return
		}

		c.Set("ctx-user-id", claims.UserId)
	}
}
