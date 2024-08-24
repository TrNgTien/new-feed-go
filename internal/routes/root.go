package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func InitialRoutes(r *gin.Engine) {
	mainGroup := r.Group("/api/v1")

	PORT := fmt.Sprintf(": %s", os.Getenv("APP_PORT"))
	r.Use(gin.Recovery())

	AuthRoutes(mainGroup)
	UserRoutes(mainGroup)
	FeedRoutes(mainGroup)

	r.Run(PORT)
}
