package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/te-fa-bene/api-go/docs"
	"github.com/te-fa-bene/api-go/internal/app/database"
	"github.com/te-fa-bene/api-go/internal/app/http/handler"
	"github.com/te-fa-bene/api-go/internal/app/http/middleware"
	"github.com/te-fa-bene/api-go/internal/app/repository"
	"github.com/te-fa-bene/api-go/internal/app/service"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	employeeRepository := repository.NewEmployeeRepository(db)
	authService := service.NewAuthService(employeeRepository)
	authHandler := handler.NewAuthHandler(authService, employeeRepository)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.Health)

		v1.POST("/auth/login", authHandler.Login)
		v1.GET("/me", middleware.Auth(), authHandler.Me)
	}

	return r
}
