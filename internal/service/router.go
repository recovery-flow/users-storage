package service

import (
	"context"
	"time"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVICE)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVICE, service))
	authMW := service.TokenManager.Middleware(service.Config.JWT.AccessToken.SecretKey)
	rateLimiter := httpkit.NewRateLimiter(15, 10*time.Second, 5*time.Minute)

	r.Route("/user-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(rateLimiter.Middleware)
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/create", handlers.CreateUser)
				r.Route("/update", func(r chi.Router) {
					r.Patch("/username", handlers.UpdateUsername)
					r.Patch("/title", handlers.UpdateTitle)
					r.Patch("/status", handlers.UpdateStatus)
					r.Patch("/avatar", handlers.UpdateAvatar)
					r.Patch("/bio", handlers.UpdateBio)
				})
			})
			r.Route("/public", func(r chi.Router) {
				r.Get("/get/{username}", handlers.GetUser)
				r.Get("/search", handlers.SearchUsers)
			})
		})
	})
}
