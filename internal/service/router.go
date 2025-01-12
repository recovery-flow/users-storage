package service

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/handlers"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMdl(service.Config.JWT.AccessToken.SecretKey)
	rateLimiter := httpkit.NewRateLimiter(service.Config.Rate.MaxRequests, service.Config.Rate.TimeWindow, service.Config.Rate.Expiration)

	r.Route("/users-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(rateLimiter.Middleware)
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Route("/user", func(r chi.Router) {
					r.Route("/update", func(r chi.Router) {
						r.Put("/", handlers.UpdateUser)
						r.Post("/avatar", handlers.UpdateAvatar)
						r.Patch("/username", handlers.UpdateUsername)
					})
				})

				r.Route("/team", func(r chi.Router) {
					r.Post("/create", handlers.CreateTeam)
					r.Route("/update", func(r chi.Router) {
						r.Put("/", handlers.UpdateTeam)
					})
					r.Route("/{team_id}", func(r chi.Router) {
						r.Route("member", func(r chi.Router) {
							r.Post("/create", handlers.CreateMember)
							r.Delete("/remove", handlers.RemoveMember)
							r.Patch("/role", handlers.UpdateMemberRole)
						})
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/user", func(r chi.Router) {
					r.Get("/get/{username}", handlers.GetUser)
				})

				r.Route("/team", func(r chi.Router) {
					r.Get("/{tram_id}", handlers.GetTeam)
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
