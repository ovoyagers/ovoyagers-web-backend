package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/routes"
	"github.com/stretchr/testify/assert"
)

var router = gin.Default()

func TestHealthRoute_HealthCheck(t *testing.T) {
	router.GET("/health/ping", routes.HealthHandler)
	req, err := http.NewRequest("GET", "/health/ping", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
