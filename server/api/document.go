package api

import (
	"net/http"
	"strconv"

	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/gin-gonic/gin"
)

type documentDatabase interface {
	CreateDocument(doc model.Document) error
	GetDocumentByID(id uint) (model.Document, error)
	DeleteDocumentByID(id uint) error
	GetDocuments() ([]model.Document, error)
}

type DocumentApi struct {
	DB documentDatabase
}

func (d *DocumentApi) GetDocument(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Could not parse id: " + err.Error()})
		return
	}

	if doc, err := d.DB.GetDocumentByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Document not found"})
	} else {
		c.JSON(http.StatusOK, doc)
	}

}

func (d *DocumentApi) GetDocuments(c *gin.Context) {

	documents, err := d.DB.GetDocuments()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not get documents: " + err.Error()})
	}
	c.JSON(http.StatusOK, documents)

}

func (d *DocumentApi) CreateDocument(c *gin.Context) {

	var newDoc model.Document

	if err := c.Bind(&newDoc); err == nil {
		if err := d.DB.CreateDocument(newDoc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not create document: " + err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Document created"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Could not create document: " + err.Error()})
	}

}

func (d *DocumentApi) DeleteDocument(c *gin.Context) {

	lookfor, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Could not parse id: " + err.Error()})
	}
	if err := d.DB.DeleteDocumentByID(uint(lookfor)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Could not delete document: " + err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Document deleted"})
	}

}
