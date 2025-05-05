// api-seguridad/core/middleware/auth.go
package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simplemente continuamos sin verificar autenticación
		// Establecemos un userID por defecto (1 para admin o 0 si prefieres)
		c.Set("userID", uint(1))
		c.Next()
	}
}