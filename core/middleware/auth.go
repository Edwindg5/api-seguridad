// api-seguridad/core/middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener token del header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authentication required",
				"error":   "Authorization header missing",
			})
			return
		}

		// Validación básica del token (en producción usar JWT real)
		if token != "Bearer valid-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authentication failed",
				"error":   "Invalid token",
			})
			return
		}

		// Simulamos extraer el ID del usuario del token
		// En producción, esto vendría del JWT decodificado
		userID := uint(1) // ID de ejemplo
		c.Set("userID", userID)
		c.Next()
	}
}