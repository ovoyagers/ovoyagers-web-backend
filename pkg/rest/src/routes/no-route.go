package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoRoute struct {
}

func NewNoRoute() *NoRoute {
	return &NoRoute{}
}

func (r *NoRoute) NoRouteFound(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}
