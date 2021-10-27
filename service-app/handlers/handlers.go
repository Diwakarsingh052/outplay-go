package handlers

import (
	"log"
	"net/http"
	"os"
	"service-app/middleware"
	"service-app/web"
)

func API(shutdown chan os.Signal, log *log.Logger) http.Handler {
	//app := mux.NewRouter()
	app := web.NewApp(shutdown)
	m := middleware.Mid{Log: log}

	app.HandleFunc(http.MethodGet, "/ready", m.Logger(m.Errors(m.Panics(check))))
	return app
}
