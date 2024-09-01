package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TrNgTien/new-feed-go/internal/datasource/cache"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	ctx := context.Background()

	err := cache.Set(ctx, "key1", "hello", 0)

	if err != nil {
		fmt.Println("[Login] err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set cache"})
		return
	}

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
