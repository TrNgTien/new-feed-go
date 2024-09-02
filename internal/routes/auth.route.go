package routes

import (
	"github.com/TrNgTien/new-feed-go/internal/constants"
	"github.com/TrNgTien/new-feed-go/internal/controller"
	"github.com/TrNgTien/new-feed-go/internal/services"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group(constants.AUTH_PATH)

	authService := services.NewAuthService()
	authController := controller.NewAuthController(authService)
	auth.POST("sign-in", authController.Login)
	auth.POST("sign-up", authController.Register)
	auth.POST("logout", authController.Logout)
}
