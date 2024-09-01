package datasource

import (
	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/TrNgTien/new-feed-go/internal/datasource/cache"
	"github.com/TrNgTien/new-feed-go/internal/datasource/minio"
	"github.com/TrNgTien/new-feed-go/internal/datasource/mysql"
)

func InitDatasource() {
	cacheConfigs := configs.Cache{
		Address: "localhost:6379",
	}

	mysql.InitMysql()
	cache.InitRedisClient(cacheConfigs)
	minio.Init()
}
