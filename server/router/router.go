package router

import (
	"net/http"

	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]model.Document, 0)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	doc := r.Group("/document")

	// Get documents
	doc.GET("", getDocumentsHandler)

	// Get document by id
	doc.GET("/:id", getDocumentHandler)

	// Create document
	doc.POST("", createDocumentHandler)

	// Delete document
	doc.DELETE("", func(c *gin.Context) {
	})

	return r
}

func getDocumentHandler(c *gin.Context) {

	if doc, ok := db[c.Param("id")]; ok {
		c.JSON(http.StatusOK, doc)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "Document not found"})
	}

}

func getDocumentsHandler(c *gin.Context) {

	for _, doc := range db {
		c.JSON(http.StatusOK, doc)
	}

}

func createDocumentHandler(c *gin.Context) {

	var newDoc model.Document

	if err := c.Bind(&newDoc); err == nil {
		db[newDoc.ID] = newDoc
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Could not create document: " + err.Error()})
	}

}

func deleteDocumentHandler(c *gin.Context) {
	if doc, ok := db[c.Param("id")]; ok {
		delete(db, doc.ID)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "Document not found"})
	}

}
