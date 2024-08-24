package internal

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/TrNgTien/new-feed-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.InitialRoutes(r)
}
