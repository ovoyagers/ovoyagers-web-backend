package authtest

// var ctx = context.Background()
// var authController = authcontroller.NewAuthController(ctx)

// var router = gin.Default()

// func TestLoginController_Login(t *testing.T) {
// 	router.POST("/login", authController.Login)

// 	// create a buffer for the wrong country code body
// 	var buff bytes.Buffer
// 	err := json.NewEncoder(&buff).Encode(map[string]interface{}{
// 		"countryCode": "in",
// 		"phone":       "+919876543210",
// 	})

// 	assert.NoError(t, err)
// 	req, err := http.NewRequest("POST", "/login", &buff)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
// }
