// api-seguridad/core/middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

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

		// Validar formato Bearer
		if !strings.HasPrefix(token, "Bearer ") || len(token) <= 7 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authentication failed",
				"error":   "Invalid token format",
			})
			return
		}

		// Aquí podrías validar el JWT si fuera necesario
		// Simulamos extracción del userID
		userID := uint(1)
		c.Set("userID", userID)

		c.Next()
	}
}