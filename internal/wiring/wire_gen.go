// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiring

import (
	"github.com/TrNgTien/new-feed-go/internal/app"
	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/TrNgTien/new-feed-go/internal/controller"
	"github.com/TrNgTien/new-feed-go/internal/datasource"
	"github.com/TrNgTien/new-feed-go/internal/services"
	"github.com/TrNgTien/new-feed-go/internal/utils"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeServer() (*app.AppServer, func(), error) {
	appServer := app.NewAppServer()
	return appServer, func() {
	}, nil
}

// wire.go:

var WireSet = wire.NewSet(app.WireSet, configs.WireSet, utils.WireSet, controller.WireSet, services.WireSet, datasource.WireSet)
