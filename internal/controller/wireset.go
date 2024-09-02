package controller

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewAuthController,
	NewFeedController,
)
