package database

import (
	"github.com/fabian-lapotre/document-api/server/model"
	"gorm.io/gorm"
)

func (d *GormDataBase) CreateDocument(doc model.Document) error {
	return d.DB.Create(&doc).Error
}

func (d *GormDataBase) GetDocumentByID(id uint) (model.Document, error) {
	var doc model.Document
	err := d.DB.First(&doc, id).Error
	return doc, err
}

func (d *GormDataBase) DeleteDocumentByID(id uint) error {
	return d.DB.Delete(&model.Document{}, id).Error
}

func (d *GormDataBase) GetDocuments() ([]model.Document, error) {
	var docs []model.Document
	err := d.DB.Find(&docs).Error
	return docs, err
}

type GormDataBase struct {
	DB *gorm.DB
}
