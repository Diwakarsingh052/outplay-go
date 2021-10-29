package handlers

import (
	"log"
	"net/http"
	"os"
	"service-app/auth"
	"service-app/data/user"
	"service-app/middleware"
	"service-app/web"
)

func API(shutdown chan os.Signal, log *log.Logger, a *auth.Auth, uDB *user.DbService) http.Handler {
	//app := mux.NewRouter()
	app := web.NewApp(shutdown)
	m := middleware.Mid{
		Log: log,
		A:   a,
	}
	uh := userHandlers{
		DbService: uDB,
		auth:      a,
	}

	app.HandleFunc(http.MethodGet, "/ready", m.Logger(m.Errors(m.Panics(m.Authenticate(m.HasRole(check, []string{auth.RoleAdmin}))))))
	app.HandleFunc(http.MethodPost, "/create", m.Logger(m.Errors(m.Panics(uh.SignUp))))
	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Errors(m.Panics(uh.Login))))
	return app
}
