package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/fabian-lapotre/document-api/server/api/mocks"
	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MessageSuite struct {
	apiToTest *DocumentApi
	ctx       *gin.Context
}

func TestDocumentApi_GetDocuments(t *testing.T) {

	//Given
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	// Oject to get
	expectedDocuments := []model.Document{}

	//Set a mock to get a new instance of the DocumentApi
	mockDb := mocks.NewDocumentDatabase(t)
	mockDb.On("GetDocuments").Return(expectedDocuments, nil)

	//When
	apiToTest := DocumentApi{DB: mockDb}
	apiToTest.GetDocuments(ctx)

	//Then

	// get reponson from api
	bytes, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	actual := string(bytes)

	// serialize expected Document to json
	jsonInBytes, err := json.Marshal(expectedDocuments)
	assert.Nil(t, err)
	epectedInObjJSON := string(jsonInBytes)

	assert.JSONEq(t, epectedInObjJSON, actual)

}

//TODO: implement other united test to test api
