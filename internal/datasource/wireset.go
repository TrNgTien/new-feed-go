package datasource

import (
	"github.com/TrNgTien/new-feed-go/internal/datasource/cache"
	"github.com/TrNgTien/new-feed-go/internal/datasource/minio"
	"github.com/TrNgTien/new-feed-go/internal/datasource/mysql"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	minio.WireSet,
	mysql.WireSet,
	cache.WireSet,
)
