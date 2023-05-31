package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Shellback.nl/Restapi/router"
	"github.com/stretchr/testify/assert"
)

func TestGetNotification(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req, _ := http.NewRequest("GET", "/getNotifications/"+"testing", nil) // add token here
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, strings.Contains(w.Body.String(), "CVE 8.1 Found"))
}
