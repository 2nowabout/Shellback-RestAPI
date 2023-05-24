package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"Shellback.nl/Restapi/router"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationMiddlewareNoAuthentication(t *testing.T) {
	// Maak een nieuwe router
	router := router.RequestHandler()

	// Registreer de middleware voor authenticatie
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getCompanies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "Unauthorized", w.Body.String())
}

func TestAuthenticationLoginInvalid(t *testing.T) {
	router := router.RequestHandler()

	// Registreer de middleware voor authenticatie
	w := httptest.NewRecorder()
	body := `{"username":"wrongusername", "password":"wrongPassword"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	fmt.Printf(w.Body.String())
}

func TestAuthenticationLoginValid(t *testing.T) {
	router := router.RequestHandler()
	err := godotenv.Load("../.env")
	username := os.Getenv("TESTINGUSERNAME")
	password := os.Getenv("TESTINGPASSWORD")
	if err != nil {
		fmt.Println("Error loading environment variables file")
	}
	// Registreer de middleware voor authenticatie
	w := httptest.NewRecorder()
	body := `{"username":"` + username + `", "password":"` + password + `"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	router.ServeHTTP(w, req)

	fmt.Printf(w.Body.String())
	assert.Equal(t, 200, w.Code)

}
