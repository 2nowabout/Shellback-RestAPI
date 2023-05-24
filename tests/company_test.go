package tests

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Shellback.nl/Restapi/router"
	"github.com/stretchr/testify/assert"
)

func TestGetCompany(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req, _ := http.NewRequest("GET", "/getCompany/"+"2", nil) // add token here
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, strings.Contains(w.Body.String(), "Testing2"))
}

func TestGetAllCompanies(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req, _ := http.NewRequest("GET", "/getCompanies", nil) // add token here
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, strings.Contains(w.Body.String(), "Testing"))
}

func TestGetCompanyNotificationAmount(t *testing.T) {
	router := router.RequestHandler()
	token := loginAndReturnToken(router)

	req, _ := http.NewRequest("GET", "/getCompany/notificationAmount/"+"Testing2", nil) // add token here
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	buf := w.Body
	var value int
	err := binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 9, value)
}
