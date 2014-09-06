package test

import (
	"net/http"
	"testing"

	"../route"
)

func TestIndex(t *testing.T) {
	Request("GET", "/", route.GoalsIndex)

	if response.Code != http.StatusOK {
		t.Errorf("Didn't return a 200 response.")
	}
}
