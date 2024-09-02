package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func R200(c *gin.Context, data any) {
	c.JSON(http.StatusOK,
		gin.H{
			"data":   data,
			"status": "success",
		})
}

func R400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"message": msg,
			"status":  "error",
		})
}

func R401(c *gin.Context) {
	c.JSON(http.StatusUnauthorized,
		gin.H{
			"message": "Unauthorized!",
			"status":  "error",
		})
}

func R403(c *gin.Context, data any) {
	c.JSON(http.StatusForbidden,
		gin.H{
			"message": "Forbidden!",
			"status":  "error",
		})
}

func R500(c *gin.Context, data any) {
	c.JSON(http.StatusInternalServerError,
		gin.H{
			"message": "Internal server error!",
			"status":  "error",
		})
}
