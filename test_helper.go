package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-martini/martini"
)

var (
	response *httptest.ResponseRecorder
)

//
// Request simulates a request for a given method, route and handler.
//
func Request(method string, route string, handler martini.Handler) {
	initDatabase()
	m := martini.Classic()
	m.Get(route, handler)

	request, _ := http.NewRequest(method, route, nil)
	response = httptest.NewRecorder()

	m.ServeHTTP(response, request)
}
