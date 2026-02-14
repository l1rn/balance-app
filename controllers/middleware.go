package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/l1rn/balance-app/services"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Auth header required"})
			return
		}

		tokenString := strings.Split(authHeader, "Bearer")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return services.SecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid Token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid Claims"})
			return
		}

		userRole := claims["role"].(string)
		if requiredRole != "" && userRole != requiredRole {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins only!"})
			return
		}

		ctx.Set("userId", claims["sub"])
		ctx.Next()
	}
}
