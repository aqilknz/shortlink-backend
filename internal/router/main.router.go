package router

import (
	"github.com/aqilknz/shortlink-backend/internal/controller"
	"github.com/aqilknz/shortlink-backend/internal/middleware"
	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/aqilknz/shortlink-backend/internal/service"
	"github.com/aqilknz/shortlink-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	_ "github.com/aqilknz/shortlink-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(app *gin.Engine, db *pgxpool.Pool, redis *redis.Client) {
	app.Use(middleware.CORSMiddleware)
	// Middleware Global
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRepo := repository.NewAuthRepository(db, redis)
	linkRepo := repository.NewLinkRepository(db)
	hashConfig := pkg.NewHashConfig()
	authService := service.NewAuthService(authRepo, hashConfig)
	linkService := service.NewLinkService(linkRepo)

	linkController := controller.NewLinkController(linkService)
	authController := controller.NewAuthController(authService)

	api := app.Group("/api")
	RegisterAuthRoutes(api, authController, authRepo)
	RegisterLinkRoutes(api, linkController, authRepo)
	RegisterRedirectRoute(app, linkController)
}
