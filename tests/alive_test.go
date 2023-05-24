package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Shellback.nl/Restapi/router"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationGetAllAlive(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req2, _ := http.NewRequest("GET", "/getAllAlive", nil) // add token here
	req2.Header.Set("Authorization", "Bearer "+token)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
}

func TestAuthenticationGetSingleAlive(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req2, _ := http.NewRequest("GET", "/getAlive/"+"1", nil)
	req2.Header.Set("Authorization", "Bearer "+token)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, true, strings.Contains(w2.Body.String(), "testing"))
}
