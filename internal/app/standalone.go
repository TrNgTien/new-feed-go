package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AppServer struct {
	Router *gin.Engine
	logger *log.Logger
}

func NewAppServer() *AppServer {
	mode := os.Getenv("APP_MODE")
	gin.SetMode(mode)
	return &AppServer{Router: gin.Default(), logger: log.Default()}
}

func (a *AppServer) Start() error {
	err := godotenv.Load()

	if err != nil {
		a.logger.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:        fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		ReadTimeout: time.Minute,
		Handler:     a.Router,
	}

	a.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.InitialRoutes(a.Router)

	return httpServer.ListenAndServe()
}
