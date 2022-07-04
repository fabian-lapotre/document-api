package api

import (
	"testing"

	"github.com/fabian-lapotre/document-api/server/api/mocks"
	"github.com/fabian-lapotre/document-api/server/model"
)

func TestDocumentApi_GetDocuments(t *testing.T) {

	//Given

	//Set a mock to get a new instance of the DocumentApi

	test := mocks.NewDocumentDatabase(t)
	test.On("GetDocuments").Return([]model.Document{}, nil)

	//When

	//Then

}
