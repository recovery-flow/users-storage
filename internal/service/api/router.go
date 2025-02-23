package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/tokens/identity"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/api/handlers"
)

func Run(ctx context.Context, svc *service.Service) {
	r := chi.NewRouter()

	h, err := handlers.NewHandlers(svc)
	if err != nil {
		svc.Log.Fatalf("failed to create handlers: %v", err)
		<-ctx.Done()
		return
	}

	authMW := tokens.AuthMdl(svc.Config.JWT.AccessToken.SecretKey)
	roleGrant := tokens.IdentityMdl(svc.Config.JWT.AccessToken.SecretKey, identity.Admin, identity.SuperUser)

	r.Route("/users-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)

				r.Route("/user", func(r chi.Router) {
					r.Put("/", h.UserUpdate)
					r.Post("/avatar", h.UserUpdateAvatar)
					r.Delete("/avatar", h.UserDeleteAvatar)
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/users", func(r chi.Router) {
					r.Route("/{user_id}", func(r chi.Router) {
						r.Get("/", h.UserGet)
					})

					r.Get("/filter", h.UsersFilter)
				})
			})

			r.Route("/admin", func(r chi.Router) {
				r.Use(roleGrant)
				r.Route("/users", func(r chi.Router) {
					r.Route("/{user_id}", func(r chi.Router) {
						r.Patch("/", h.AdminUserUpdate)
						r.Delete("/avatar", h.AdminDeleteAvatar)
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, svc.Config.Server.Port, r, svc.Log)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, svc.Log)
}
