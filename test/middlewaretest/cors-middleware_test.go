package middlewaretest

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

func TestCorsMiddlewareWithCustomDomain(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware("https://example.com"))
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Errorf("Expected Access-Control-Allow-Origin to be 'https://example.com', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestCorsMiddlewareWithNoDomain(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware(""))
	router.GET("/health/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health/ping", nil)
	router.ServeHTTP(w, req)
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin to be '*', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestCorsMiddlewareWithRequestMethodOptions(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware(""))
	router.OPTIONS("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("OPTIONS", "/test", nil)
	router.ServeHTTP(w, req)
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin to be '*', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}
}
