package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-martini/martini"

	"../model"
)

var (
	response *httptest.ResponseRecorder
)

//
// Request simulates a request for a given method, route and handler.
//
func Request(method string, route string, handler martini.Handler) {
	model.InitDatabase()
	m := martini.Classic()
	m.Get(route, handler)

	request, _ := http.NewRequest(method, route, nil)
	response = httptest.NewRecorder()

	m.ServeHTTP(response, request)
}
