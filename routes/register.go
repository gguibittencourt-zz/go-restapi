package routes

import (
	"context"
	"github.com/gguibittencourt/go-restapi/handler/tasks"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Params struct {
	fx.In

	Logger    *zap.Logger
	Lifecycle fx.Lifecycle
	Handler   tasks.Handler
}

func Register(p Params) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Get("/tasks", p.Handler.List)
	router.Get("/tasks/{id}", p.Handler.Find)
	router.Post("/tasks", p.Handler.Create)
	router.Put("/tasks/{id}", p.Handler.Update)
	router.Delete("/tasks/{id}", p.Handler.Delete)
	server := http.Server{
		Addr:    ":3001",
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
