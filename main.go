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
		handler.Module,
		loggerfx.Module,
		database.Module,
		fx.Invoke(routes.Register),
	)
}
