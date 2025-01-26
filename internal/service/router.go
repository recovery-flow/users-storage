package service

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/roles"
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
	adminGrant := service.TokenManager.RoleGrant(service.Config.JWT.AccessToken.SecretKey, string(roles.RoleUserAdmin))

	r.Route("/users-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)

				r.Route("/user", func(r chi.Router) {
					r.Route("/update", func(r chi.Router) {
						r.Put("/{user_id}", handlers.UserUpdate)
						r.Post("/avatar", handlers.UserUpdateAvatar)
						r.Delete("/avatar", handlers.UserDeleteAvatar)
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/users", func(r chi.Router) {
					r.Get("/search", handlers.UsersSearch)
					r.Get("/filter", handlers.UsersFilter)
					r.Route("/{user_id}", func(r chi.Router) {
						r.Get("/", handlers.UserGet)
					})
				})
			})

			r.Route("/admin", func(r chi.Router) {
				r.Use(adminGrant)
				r.Route("/users", func(r chi.Router) {
					r.Route("/{user_id}", func(r chi.Router) {
						r.Patch("/", handlers.AdminUserUpdate)
						r.Delete("/avatar", handlers.AdminDeleteAvatar)
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
