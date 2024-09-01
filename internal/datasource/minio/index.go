package minio

import (
	"github.com/TrNgTien/new-feed-go/internal/utils"
	min "github.com/minio/minio-go"
)

var client *min.Client

func Init() {

	var (
		err error
	)

	client, err = min.New(utils.GetValueEnv("APP_ENV_MINIO_ENDPOINT", "localhost:9000"),
		utils.GetValueEnv("APP_ENV_MINIO_ACCESS_KEY", "minio-root"),
		utils.GetValueEnv("APP_ENV_MINIO_SECRET_KEY", "minio-secret-key"),
		false)

	if err != nil {
		panic(err)
	}
}

func GetClient() *min.Client {
	return client
}
