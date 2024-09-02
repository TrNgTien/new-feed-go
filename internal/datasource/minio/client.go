package minio

import (
	"fmt"

	"github.com/TrNgTien/new-feed-go/internal/utils"
	min "github.com/minio/minio-go"
)

func NewMinioClient() *min.Client {
	fmt.Println("[InitMysql] Connecting Minio!!")

	client, err := min.New(utils.GetValueEnv("APP_ENV_MINIO_ENDPOINT", "localhost:9000"),
		utils.GetValueEnv("APP_ENV_MINIO_ACCESS_KEY", "minio-root"),
		utils.GetValueEnv("APP_ENV_MINIO_SECRET_KEY", "minio-secret-key"),
		false)

	if err != nil {
		panic(err)
	}

	fmt.Println("[InitMysql] Connected Minio!!")
	return client

}
