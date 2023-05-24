package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Shellback.nl/Restapi/router"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTypes(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req, _ := http.NewRequest("GET", "/getTypes", nil) // add token here
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, strings.Contains(w.Body.String(), "1"))
	assert.Equal(t, true, strings.Contains(w.Body.String(), "2"))
	assert.Equal(t, true, strings.Contains(w.Body.String(), "3"))
	assert.Equal(t, true, strings.Contains(w.Body.String(), "4"))
	assert.Equal(t, true, strings.Contains(w.Body.String(), "5"))
}
