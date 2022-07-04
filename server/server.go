package server

import (
	"log"
	"os"

	"github.com/fabian-lapotre/document-api/server/database"
	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/fabian-lapotre/document-api/server/router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new router
func Create() *gin.Engine {

	databasePath := os.ExpandEnv("${DATABASE_LOCALIZATION}/${DATABASE_NAME}")

	log.Printf("creating database inside %v ", databasePath)

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Document{})
	if err != nil {
		panic(err)
	}

	return router.SetupRouter(&database.GormDataBase{DB: db})

}
