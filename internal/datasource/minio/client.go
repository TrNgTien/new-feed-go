package minio

import (
	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/TrNgTien/new-feed-go/internal/utils"
	min "github.com/minio/minio-go"
	"go.uber.org/zap"
)

func NewMinioClient(minioConfig configs.Storage, logger *zap.Logger) (*min.Client, func(), error) {
	logger.Info("[InitMinio] Connecting to Minio")

	client, err := min.New(utils.GetValueEnv("APP_ENV_MINIO_ENDPOINT", "localhost:9000"),
		utils.GetValueEnv("APP_ENV_MINIO_ACCESS_KEY", "minio-root"),
		utils.GetValueEnv("APP_ENV_MINIO_SECRET_KEY", "minio-secret-key"),
		false)

	if err != nil {
		logger.Error("[InitMinio] Error connecting to Minio", zap.Error(err))
		return nil, nil, err
	}

	logger.Info("[InitMinio] Connected to Minio")

	cleanup := func() {
		logger.Info("[InitMinio] Cleanup function called")
	}

	return client, cleanup, nil
}
