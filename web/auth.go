package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func EnsureAuthenticatedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		tokenString := authorization[7:]

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})

			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("AUTH_KEY"), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})

			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Expired token. Please login again"})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
