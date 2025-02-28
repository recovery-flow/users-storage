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

	r.Use(
		httpkit.CtxMiddleWare(
			handlers.CtxLog(svc.Log),
			handlers.CtxDomain(svc.Domain),
			handlers.CtxConfig(svc.Config),
		),
	)

	authMW := tokens.AuthMdl(svc.Config.JWT.AccessToken.SecretKey)
	roleGrant := tokens.IdentityMdl(svc.Config.JWT.AccessToken.SecretKey, identity.Admin, identity.SuperUser)

	r.Route("/users-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)

				r.Route("/user", func(r chi.Router) {
					r.Put("/", handlers.UserUpdate)
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/users", func(r chi.Router) {
					r.Route("/{user_id}", func(r chi.Router) {
						r.Get("/", handlers.UserGet)
					})

					r.Get("/filter", handlers.UsersFilter)
				})
			})

			r.Route("/admin", func(r chi.Router) {
				r.Use(roleGrant)
				r.Route("/users", func(r chi.Router) {
					r.Route("/{user_id}", func(r chi.Router) {
						r.Patch("/", handlers.AdminUserUpdate)
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, svc.Config.Server.Port, r, svc.Log)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, svc.Log)
}
