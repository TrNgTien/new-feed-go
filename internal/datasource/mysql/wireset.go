package mysql

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewMysqlClient,
)
