package routes

import (
	"net/http"

	"github.com/TrNgTien/new-feed-go/internal/constants"
	"github.com/gin-gonic/gin"
)

func FeedRoutes(rg * gin.RouterGroup) {
	feeds := rg.Group(constants.FEED_PATH)

	feeds.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "feed")
	})

	feeds.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"data": "feed id",
			})
	})

}
