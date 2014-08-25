package main

import (
	"net/http"
	"testing"
)

func TestIndex(t *testing.T) {
	Request("GET", "/", HandleIndex)

	if response.Code != http.StatusOK {
		t.Errorf("Didn't return a 200 response.")
	}
}
