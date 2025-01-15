package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
	"github.com/petmeds24/backend/pkg/rest/src/models"
)

type HealthRoute struct {
}

// NewHealthRoute is a constructor function that creates a new instance of the HealthRoute struct.
//
// It takes no parameters.
//
// It returns a pointer to a HealthRoute struct.
func NewHealthRoute() *HealthRoute {
	// Create a new instance of the HealthRoute struct and return a pointer to it.
	return &HealthRoute{}
}

func (hr *HealthRoute) SetupHealthRoute(rg *gin.RouterGroup) {
	healthRoute := rg.Group("/health")
	{
		healthRoute.GET("/ping", HealthHandler)
		healthRoute.GET("/status", middlewares.DeserializeUser(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, models.Response{
				Message:    "private route",
				Data:       nil,
				Status:     "success",
				StatusCode: http.StatusOK,
			})
		})
	}
}

// HealthHandler godoc
//
//	@Summary		Health check
//	@Description	Health check
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Router			/health/ping [get]
func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Message:    "server is up and running",
		Data:       nil,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}

// PrivateTestHandler godoc
//
//	@Summary		Private route testing
//	@Description	Private route testing
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Router			/health/status [get]
//	@Security		BearerAuth
func PrivateTestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Message:    "private route testing",
		Data:       nil,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}
