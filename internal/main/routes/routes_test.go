package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	_ "main/internal/config/env"

	"github.com/stretchr/testify/assert"
)

// Not working yet

// Test register new username
func TestRegister(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("email", "pras_iman@hotmail.com")
	form.Add("password", "pras_iman@hotmail.com")
	req, _ := http.NewRequest("POST", "/api/v1/register", strings.NewReader(form.Encode()))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Test with nonexisted credentials
func TestLogin(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("email", "myusername")
	form.Add("password", "mypassword")
	req, _ := http.NewRequest("POST", "/api/v1/auth", strings.NewReader(form.Encode()))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
