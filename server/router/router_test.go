package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fabian-lapotre/document-api/server/database"
	"github.com/fabian-lapotre/document-api/server/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	client = &http.Client{}
	db     *gorm.DB
)

type IntegrationSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupTest() {

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Document{})
	db.Create(&model.Document{ID: 1, Name: "Document 1", Description: "Content 1"})
	db.Create(&model.Document{ID: 2, Name: "Document 2", Description: "Content 2"})
	db.Create(&model.Document{ID: 3, Name: "Document 3", Description: "Content 3"})

	s.server = httptest.NewServer(SetupRouter(&database.GormDataBase{DB: db}))
}

func (s *IntegrationSuite) TearDownTest() {
	os.Remove("./test.db")
	s.server.Close()
}

func (s *IntegrationSuite) TestAddDocument() {
	// Add a new document and test everything is ok
	req := s.newRequest("POST", "document", `{"id":4,"name":"Document 4", "description":"Content 4"}`)
	doRequestAndExpect(s.T(), req, 200, `{"status": "Document created"}`)

	// Add a document with an existing id and test error
	req = s.newRequest("POST", "document", `{"id":1,"name":"Document 1", "description":"Content 1"}`)
	doRequestAndExpect(s.T(), req, 409, `{"status": "Document already exists"}`)

	// Send bad json and test error
	req = s.newRequest("POST", "document", `{"id":4,"nane":"Document 4", "description":"Content 4"}`)
	doRequestAndExpect(s.T(), req, 400, `{"status":"Could not create document: Key: 'Document.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`)
}

func (s *IntegrationSuite) TestGetDocument() {
	req := s.newRequest("GET", "document/1", "")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var firstDocumentReceived model.Document
	var firstDocumentSaveInDB model.Document

	json.NewDecoder(res.Body).Decode(&firstDocumentReceived)

	db.First(&firstDocumentSaveInDB, 1)

	assert.Equal(s.T(), firstDocumentSaveInDB, firstDocumentReceived)

	var secondDocumentReceived model.Document
	var secondDocumentSaveInDB model.Document

	req = s.newRequest("GET", "document/2", "")
	resSecondDocument, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resSecondDocument.Body).Decode(&secondDocumentReceived)
	db.First(&secondDocumentSaveInDB, 2)

	assert.Equal(s.T(), secondDocumentSaveInDB, secondDocumentReceived)

	req = s.newRequest("GET", "document/3", "")
	resThirdDocument, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var thirdDocumentReceived model.Document
	var thirdDocumentSaveInDB model.Document

	json.NewDecoder(resThirdDocument.Body).Decode(&thirdDocumentReceived)
	db.First(&thirdDocumentSaveInDB, 3)

	assert.Equal(s.T(), thirdDocumentSaveInDB, thirdDocumentReceived)

	req = s.newRequest("GET", "document/40", "")
	doRequestAndExpect(s.T(), req, 404, `{"status": "Document not found"}`)

}

func (s *IntegrationSuite) TestDeleteDocument() {
	req := s.newRequest("DELETE", "document/1", "")
	doRequestAndExpect(s.T(), req, 200, `{"status": "Document deleted"}`)

	// Delete a document that does not exist
	// gorm all to delete element that doesn exist need to investigate on it
	/*
		req = s.newRequest("DELETE", "document/100", "")
		doRequestAndExpect(s.T(), req, 404, `{"status": "Document not found"}`)
	*/

}

func (s *IntegrationSuite) TestGetDocuments() {

	req := s.newRequest("GET", "document", "")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// documentRec
	var documentsReceived []model.Document
	json.NewDecoder(res.Body).Decode(&documentsReceived)

	var documentsSaveInDB []model.Document
	result := db.Find(&documentsSaveInDB)

	fmt.Println(result)

	assert.Equal(s.T(), documentsSaveInDB, documentsReceived)

}

func (s *IntegrationSuite) newRequest(method, url, body string) *http.Request {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", s.server.URL, url), strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	assert.Nil(s.T(), err)
	return req
}

func doRequestAndExpect(t *testing.T, req *http.Request, code int, json string) {
	res, err := client.Do(req)
	assert.Nil(t, err)
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	assert.Equal(t, code, res.StatusCode)
	assert.JSONEq(t, json, buf.String())
}
