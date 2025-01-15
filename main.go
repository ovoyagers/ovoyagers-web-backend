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
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Set up configuration
	globalCfg := config.NewGlobalConfig(ctx)
	cfg := globalCfg.GetConfig()

	baseURL, schema := GetBaseURLAndSchema(cfg)
	UpdateSwaggerInfo(baseURL, schema)

	// Set up Gin server
	server := gin.Default()
	server.Use(middlewares.CorsMiddleware("*"))

	// Set up routes
	router := server.Group("/api/v1")
	mainRoute := routes.NewMainRoute(globalCfg, router)
	mainRoute.SetupRoutes()

	// Swagger configuration
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true), ginSwagger.URL(baseURL)))

	// No route found
	noRoute := routes.NewNoRoute()
	server.NoRoute(noRoute.NoRouteFound)

	// Log and start server
	log.Infof("Server is running at %s://%s in %s environment", schema, baseURL, cfg.ENVIRONMENT)
	go func() {
		if err := server.Run(":" + cfg.PORT); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Info("Shutting down gracefully...")
}

func UpdateSwaggerInfo(baseURL, schema string) {
	docs.SwaggerInfo.Host = baseURL
	docs.SwaggerInfo.Schemes = []string{schema}
}

func GetBaseURLAndSchema(cfg *config.Config) (string, string) {
    if cfg.ENVIRONMENT == "local" {
        return "http://localhost:4000/api/v1", "http"
    }
    return "https://ovoyagers-web-backend.onrender.com/api/v1", "https"
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
