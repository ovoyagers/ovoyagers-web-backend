package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/petmeds24/backend/config"
	docs "github.com/petmeds24/backend/docs"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
	"github.com/petmeds24/backend/pkg/rest/src/routes"
	"github.com/petmeds24/backend/pkg/rest/src/utils/constants"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	log.SetFormatter(
		&log.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// @title						Ovoyagers API
// @version					0.10
// @description				This is a backend server for Ovoyagers.
// @termsOfService				http://swagger.io/terms/
// @contact.name				Pet Care
// @contact.url				https://google.com
// @contact.email				developer@ovoyagers.com
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath					/api/v1
// @securityDefinitions.apiKey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer " before your access token
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Set up configuration
	globalCfg := config.NewGlobalConfig(ctx)
	cfg := globalCfg.GetConfig()
	consts := constants.GetConstants(cfg)

	// Set up Gin server
	server := gin.Default()
	server.Use(middlewares.CorsMiddleware())

	// Set up routes
	router := server.Group("/api/v1")
	mainRoute := routes.NewMainRoute(globalCfg, router)
	mainRoute.SetupRoutes()

	// dynamic docs
	docs.SwaggerInfo.Host = consts.BASE_URL
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{consts.SCHEMA}

	// Swagger configuration
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	// No route found
	noRoute := routes.NewNoRoute()
	server.NoRoute(noRoute.NoRouteFound)

	// Log and start server
	log.Infof("Server is running in %s environment", cfg.ENVIRONMENT)
	go func() {
		if err := server.Run(":" + cfg.PORT); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Info("Shutting down gracefully...")
}
