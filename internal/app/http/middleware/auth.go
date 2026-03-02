package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	CtxEmployeeID = "employee_id"
	CtxStoreID    = "store_id"
	CtxRole       = "role"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Bearer token",
			})
			return
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET not set"})
			return
		}

		raw := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || token == nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		employeeID, _ := claims["sub"].(string)
		storeID, _ := claims["store_id"].(string)
		role, _ := claims["role"].(string)

		if employeeID == "" || storeID == "" || role == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			return
		}

		ctx.Set(CtxEmployeeID, employeeID)
		ctx.Set(CtxStoreID, storeID)
		ctx.Set(CtxRole, role)

		ctx.Next()
	}
}
