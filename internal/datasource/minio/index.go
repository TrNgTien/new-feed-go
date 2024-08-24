package minio

import (
	"github.com/TrNgTien/new-feed-go/internal/utils"
	min "github.com/minio/minio-go"
)

type MinioClient struct {
	client *min.Client
}

func (c *MinioClient) Init() {
	var (
		err error
	)

	c.client, err = min.New(utils.GetValueEnv("APP_ENV_MINIO_ENDPOINT", "localhost:9000"),
		utils.GetValueEnv("APP_ENV_MINIO_ACCESS_KEY", "minio-root"),
		utils.GetValueEnv("APP_ENV_MINIO_SECRET_KEY", "minio-secret-key"),
		false)

	if err != nil {
		panic(err)
	}
}

func (c *MinioClient) GetClient() *min.Client {
	return c.client
}
