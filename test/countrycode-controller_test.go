package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petmeds24/backend/pkg/rest/src/controllers"
	"github.com/stretchr/testify/assert"
)

var countryCodeController = controllers.NewCountryCodeController()

func TestNewCountryCodeRoute_GetCountryCodes(t *testing.T) {
	router.GET("/countries/all", countryCodeController.GetCountryCodes)

	req, err := http.NewRequest("GET", "/countries/all", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
