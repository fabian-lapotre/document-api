package server

import (
	"github.com/fabian-lapotre/document-api/server/database"
	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/fabian-lapotre/document-api/server/router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new router
func Create(initialDb map[string]model.Document) *gin.Engine {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return router.SetupRouter(&database.GormDataBase{DB: db})

}
