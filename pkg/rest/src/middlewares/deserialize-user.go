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
		authHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authHeader)
		if len(fields) != 2 || fields[0] != "Bearer" || fields[1] == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
				Message:    "no token provided",
				Error:      "Status Unauthorized",
				Status:     "error",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}
		accessToken := fields[1]
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
		ctx.Set("user_id", claims["id"].(string))
		ctx.Next()
	}
}
