package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// Status code test
	if w.Code != 200 {
		t.Error("Http test fail")
	}

	// Return value test
	if w.Body.String() != "pong" {
		t.Error("wrong content")
	}
}
