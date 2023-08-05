package main

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
