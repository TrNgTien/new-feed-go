package routes

import (
	"net/http"

	"github.com/TrNgTien/new-feed-go/internal/constants"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group(constants.USER_PATH)

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"data": "users",
			})
	})
}
