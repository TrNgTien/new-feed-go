package internal

import (
	"log"

	"github.com/TrNgTien/new-feed-go/internal/datasource"
	"github.com/TrNgTien/new-feed-go/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	r := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	datasource.InitDatasource()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.InitialRoutes(r)
}
