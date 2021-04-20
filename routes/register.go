package routes

import (
	"context"
	"github.com/gguibittencourt/go-restapi/handler/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Params struct {
	fx.In

	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Handler    users.Handler
}

func Register(p Params) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Get("/users", p.Handler.List)
	router.Get("/users/{id}", p.Handler.Find)
	router.Post("/users", p.Handler.Create)
	router.Put("/users/{id}", p.Handler.Update)
	router.Delete("/users/{id}", p.Handler.Delete)
	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				p.Logger.Info("Starting server.")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				p.Logger.Info("Shutting down server.")
				return server.Shutdown(ctx)
			},
		},
	)
}
