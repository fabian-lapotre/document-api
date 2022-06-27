package router

import (
	"net/http"
	_ "net/http/httptest"
)

var (
	client = &http.Client{}
)
