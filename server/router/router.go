package router

import (
	"net/http"

	"github.com/fabian-lapotre/document-api/server/api"
	"github.com/fabian-lapotre/document-api/server/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter(initialDb *database.GormDataBase) *gin.Engine {

	documentHandler := api.DocumentApi{DB: initialDb}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	doc := r.Group("/document")

	// Get documents
	doc.GET("", documentHandler.GetDocuments)

	// Get document by id
	doc.GET("/:id", documentHandler.GetDocument)

	// Create document
	doc.POST("", documentHandler.CreateDocument)

	// Delete document
	doc.DELETE("/:id", documentHandler.DeleteDocument)

	return r
}
