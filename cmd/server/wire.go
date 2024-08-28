//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tpl-x/conec/internal/config"
	"github.com/tpl-x/conec/internal/server"
	"github.com/tpl-x/conec/internal/service"
	"github.com/tpl-x/conec/pkg/logger"
)

var appSet = wire.NewSet(
	config.LoadDefaultConfig,
	wire.FieldsOf(new(*config.ServerConfig), "Log"),
	logger.NewZapLogger,
)

// wireApp init for builder backend
func wireApp() (*app, error) {
	panic(
		wire.Build(
			appSet,
			service.ProviderSet,
			server.ProviderSet,
			newApp))
}
