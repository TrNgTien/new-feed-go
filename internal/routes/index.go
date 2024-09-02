package routes

import (
	"github.com/gin-gonic/gin"
)

func InitialRoutes(r *gin.Engine) {
	mainGroup := r.Group("/api/v1")

	r.Use(gin.Recovery())

	AuthRoutes(mainGroup)
	UserRoutes(mainGroup)
	FeedRoutes(mainGroup)

}
