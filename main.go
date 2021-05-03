package main

import (
	"github.com/gguibittencourt/go-restapi/handler"
	"github.com/gguibittencourt/go-restapi/modules/database"
	"github.com/gguibittencourt/go-restapi/modules/loggerfx"
	"github.com/gguibittencourt/go-restapi/routes"
	"go.uber.org/fx"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		loggerfx.Module,
		database.Module,
		handler.Module,
		fx.Invoke(routes.Register),
	)
}
