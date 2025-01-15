package middlewaretest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeUser_WithToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Endpoint to generate token
	router.GET("/token", func(ctx *gin.Context) {
		jwtUtil := utils.NewJWTUtil()
		token, err := jwtUtil.CreateToken("123", "test@example.com")
		assert.NoError(t, err)
		ctx.JSON(http.StatusOK, token)
	})

	// Protected endpoint requiring token
	router.GET("/test", middlewares.DeserializeUser(), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	// Request to get the token
	req, err := http.NewRequest(http.MethodGet, "/token", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse token from response
	var response *utils.TokenDetails
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	token := response.AccessToken

	// Request to protected endpoint with token
	req, err = http.NewRequest(http.MethodGet, "/test", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check if the request was successful
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeserializeUser_WithInvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Protected endpoint requiring token
	router.GET("/test", middlewares.DeserializeUser(), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	// Request to protected endpoint with token
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer invalid_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check if the request was successful
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeserializeUser_WithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/test", middlewares.DeserializeUser(), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}

	var errResp models.Error
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if errResp.Message != "no token provided" {
		t.Errorf("expected error message 'no token provided', got %s", errResp.Message)
	}
}
