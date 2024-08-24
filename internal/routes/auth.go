package routes

import (
	"github.com/TrNgTien/new-feed-go/internal/constants"
	"github.com/TrNgTien/new-feed-go/internal/services"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group(constants.AUTH_PATH)

	auth.POST("login", services.Login)
	auth.POST("register", services.Register)
	auth.POST("logout", services.Logout)
}
