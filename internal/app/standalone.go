package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/TrNgTien/new-feed-go/internal/datasource/cache"
	"github.com/TrNgTien/new-feed-go/internal/datasource/minio"
	"github.com/TrNgTien/new-feed-go/internal/datasource/mysql"
	"github.com/TrNgTien/new-feed-go/internal/routes"
	"github.com/TrNgTien/new-feed-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type AppServer struct {
	Router *gin.Engine
	logger *zap.Logger
}

func NewAppServer() *AppServer {
	mode := os.Getenv("APP_MODE")
	gin.SetMode(mode)
	return &AppServer{Router: gin.Default(), logger: zap.NewNop()}
}

func (a *AppServer) Start() error {
	err := godotenv.Load()

	if err != nil {
		a.logger.Fatal(err.Error())
	}

	config, err := configs.NewConfig("")
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	a.logger, _, err = utils.NewLogger(config.Log)
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	_, mysqlCleanup, err := mysql.NewMysqlClient(config.Database, a.logger)
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	defer mysqlCleanup()

	_, redisCleanup, err := cache.NewRedisClient(config.Cache, a.logger)
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	defer redisCleanup()

	_, minioCleanup, err := minio.NewMinioClient(config.Storage, a.logger)
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	defer minioCleanup()

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
