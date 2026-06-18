package router

import (
	"github.com/aqilknz/shortlink-backend/internal/controller"
	"github.com/aqilknz/shortlink-backend/internal/middleware"
	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// RegisterLinkRoutes mengatur endpoint /api/links (Membutuhkan Auth)
func RegisterLinkRoutes(rg *gin.RouterGroup, linkController *controller.LinkController, authRepo repository.AuthRepository) {
	linkGroup := rg.Group("/links")

	linkGroup.Use(middleware.RequireAuth(authRepo))
	linkGroup.POST("", linkController.CreateLink)
	linkGroup.GET("", linkController.GetUserLinks)
	linkGroup.DELETE("/:id", linkController.DeleteLink)
	linkGroup.GET("/check-slug", linkController.CheckSlugAvailability)

}
func RegisterRedirectRoute(engine *gin.Engine, linkController *controller.LinkController) {
	engine.GET("/:slug", linkController.Redirect)
}
