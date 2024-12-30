package service

import (
	"context"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/handlers"
	"github.com/go-chi/chi/v5"
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
	rateLimiter := httpkit.NewRateLimiter(service.Config.Rate.MaxRequests, service.Config.Rate.TimeWindow, service.Config.Rate.Expiration)

	r.Route("/user-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(rateLimiter.Middleware)
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Post("/create", handlers.CreateUser)
				r.Route("/update", func(r chi.Router) {
					r.Patch("/", handlers.UpdateUserFull)
					r.Patch("/username", handlers.UpdateUsername)
					r.Patch("/title", handlers.UpdateTitle)
					r.Patch("/status", handlers.UpdateStatus)
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
