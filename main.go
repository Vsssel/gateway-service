package main

import (
	"gateway/client"
	"gateway/config"
	"gateway/handlers"
	"gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gateway API
// @version 1.0
// @description API Gateway for microservices architecture
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	client.InitGRPCClients(cfg)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
    swaggerFiles.Handler,
    ginSwagger.URL("/swagger/doc.json"),
    ginSwagger.DocExpansion("none"),
    ginSwagger.DefaultModelsExpandDepth(-1),
    ginSwagger.PersistAuthorization(true),
))

	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routes := r.Routes()
		for _, route := range routes {
    	log.Printf("%s %s\n", route.Method, route.Path)
		}
	// Запуск сервера
	r.Run(cfg.Port)
}