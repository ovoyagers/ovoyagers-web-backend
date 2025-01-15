package test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/petmeds24/backend/pkg/rest/src/controllers"
// 	"github.com/stretchr/testify/assert"
// )

// var ctx = context.Background()
// var itemController = controllers.NewItemController(ctx)

// // var itemDao = daos.NewItemDao(ctx)
// var router = gin.Default()

// func TestItemController_CreateItem(t *testing.T) {
// 	router.POST("/item/create", itemController.CreateItem)

// 	// create a buffer for the response body
// 	var buff bytes.Buffer
// 	err := json.NewEncoder(&buff).Encode(map[string]interface{}{
// 		"id":   "888",
// 		"name": "unit test item",
// 	})
// 	assert.NoError(t, err)

// 	// create a request
// 	req, err1 := http.NewRequest("POST", "/item/create", &buff)
// 	assert.NoError(t, err1)

// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)

// 	// check the status code
// 	assert.Equal(t, http.StatusCreated, rec.Code)
// }

// func TestItemController_GetItem(t *testing.T) {
// 	router.GET("/item/getItem/:id", itemController.GetItem)
// 	req, err := http.NewRequest("GET", "/item/getItem/888", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// }

// func TestItemController_GetItems(t *testing.T) {
// 	router.GET("/item/getItems", itemController.GetItems)
// 	req, err := http.NewRequest("GET", "/item/getItems", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// }

// func TestItemController_UpdateItem(t *testing.T) {
// 	router.PUT("/item/updateItem", itemController.UpdateItem)

// 	// create a buffer for the response body
// 	var buff bytes.Buffer
// 	err := json.NewEncoder(&buff).Encode(map[string]interface{}{
// 		"id":   "888",
// 		"name": "unit test item",
// 	})
// 	assert.NoError(t, err)

// 	req, err := http.NewRequest("PUT", "/item/updateItem", &buff)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// }

// func TestItemController_DeleteItem(t *testing.T) {
// 	router.DELETE("/item/deleteItem/:id", itemController.DeleteItem)

// 	req, err := http.NewRequest("DELETE", "/item/deleteItem/888", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusNoContent, rec.Code)
// }
