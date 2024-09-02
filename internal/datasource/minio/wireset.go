package minio

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewMinioClient,
)
