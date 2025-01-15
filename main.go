package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	// Swagger imports
	"github.com/petmeds24/backend/config"
	_ "github.com/petmeds24/backend/docs"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
	"github.com/petmeds24/backend/pkg/rest/src/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						Pet Care API
// @version					0.01
// @description				This is a backend server for Pet Care.
// @termsOfService				http://swagger.io/terms/
// @contact.name				Pet Care
// @contact.url				https://google.com
// @contact.email				azharuddinmohammed998@gmail.com
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:4000
// @BasePath					/api/v1
// @schemes					http https
// @securityDefinitions.apiKey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer " before your access token
func main() {
	// Set up context
	ctx := context.Background()

	// Set up configuration
	globalCfg := config.NewGlobalConfig(ctx)
	cfg := globalCfg.GetConfig()

	// Set up Gin server
	server := gin.Default()

	// Middlewares
	server.Use(middlewares.CorsMiddleware("*"))

	// Set up routes
	router := server.Group("/api/v1")

	// Initialize the main route
	mainRoute := routes.NewMainRoute(globalCfg, router)

	// Swagger configuration to serve the API docs
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	// No route found
	noRoute := routes.NewNoRoute()
	server.NoRoute(noRoute.NoRouteFound)

	// Set up routes
	mainRoute.SetupRoutes()
	log.Infof("Server is running on port %s with %s environment", cfg.PORT, cfg.ENVIRONMENT)
	// Start the server and handle potential errors
	if err := server.Run(":" + cfg.PORT); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

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
