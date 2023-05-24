package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getToken(t string) string {
	token := ""
	if len(strings.Split(t, ":")) == 2 {
		token = strings.Split(t, ":")[1]
	}
	token = token[:len(token)-2]
	token = token[1:]
	return token
}

func loginAndReturnToken(router *gin.Engine) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading environment variables file")
	}
	username := os.Getenv("TESTINGUSERNAME")
	password := os.Getenv("TESTINGPASSWORD")
	// Registreer de middleware voor authenticatie
	w := httptest.NewRecorder()
	body := `{"username":"` + username + `", "password":"` + password + `"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	router.ServeHTTP(w, req)
	return getToken(strings.TrimSpace(w.Body.String()))
}
