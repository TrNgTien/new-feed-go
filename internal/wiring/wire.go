//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"github.com/TrNgTien/new-feed-go/internal/app"
	"github.com/TrNgTien/new-feed-go/internal/controller"
	"github.com/TrNgTien/new-feed-go/internal/datasource"
	"github.com/TrNgTien/new-feed-go/internal/services"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	app.WireSet,
	controller.WireSet,
	services.WireSet,
	datasource.WireSet,
)

func InitializeServer() (*app.AppServer, error) {
	wire.Build(
		WireSet,
	)

	return nil, nil
}
