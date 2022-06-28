package main

import (
	"github.com/fabian-lapotre/document-api/server"
	"github.com/fabian-lapotre/document-api/server/model"
)

func main() {

	myServer := server.Create(map[string]model.Document{})

	myServer.Run(":8080")

}
