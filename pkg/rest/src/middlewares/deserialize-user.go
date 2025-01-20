package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := utils.NewJWTUtil()

		// Get token from cookie or Authorization header
		var accessToken string

		// Try to get the token from the cookie
		cookie, err := ctx.Request.Cookie("access_token")
		if err == nil { // Only use cookie if it exists
			accessToken = cookie.Value
		}

		// If no cookie token, fall back to Authorization header
		if accessToken == "" {
			authHeader := ctx.GetHeader("Authorization")
			fields := strings.Fields(authHeader)
			if len(fields) == 2 && fields[0] == "Bearer" {
				accessToken = fields[1]
			}
		}

		// If no token is found
		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
				Message:    "no token provided",
				Error:      "Status Unauthorized",
				Status:     "error",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		// Validate the token
		claims, err := jwt.ValidateAccessToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
				Message:    err.Error(),
				Error:      "Status Unauthorized",
				Status:     "error",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		// Store user ID in the context
		if userID, ok := claims["id"].(string); ok {
			ctx.Set("user_id", userID)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{
				Message:    "invalid token payload",
				Error:      "Status Internal Server Error",
				Status:     "error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		ctx.Next()
	}
}
