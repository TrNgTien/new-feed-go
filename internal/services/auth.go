package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"data": "login",
		})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"data": "Register",
		})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"data": "Logout",
		})
}
