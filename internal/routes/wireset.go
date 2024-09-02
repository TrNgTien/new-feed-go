package routes

import "github.com/google/wire"

var WireSet = wire.NewSet(
	InitialRoutes,
)
