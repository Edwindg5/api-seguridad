// api-seguridad/core/uploads/uploads.go
package uploads

import (

	"path/filepath"
	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(router *gin.Engine) {
	// Servir archivos estáticos desde la carpeta uploads
	uploadsPath := filepath.Join("core", "uploads")
	router.Static("/uploads", uploadsPath)
}