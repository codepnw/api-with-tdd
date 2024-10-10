package test

import (
	"api-with-tdd/config"
	"api-with-tdd/routes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	app *gin.Engine
)

func TestMain(m *testing.M) {
	db = config.InitDatabase("host=localhost user=postgres password=mypassword port=4444 dbname=api_with_tdd")
	app = routes.InitRoutes(db)

	m.Run()
}

func TestHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	app.ServeHTTP(w, req)

	response := w.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["message"])
}
