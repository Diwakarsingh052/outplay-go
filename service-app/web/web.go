package web

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"syscall"
	"time"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*mux.Router // embedding router
	shutdown    chan os.Signal
}

func NewApp(shutdown chan os.Signal) *App {
	return &App{
		Router:   mux.NewRouter(),
		shutdown: shutdown,
	}
}

// HandleFunc is custom implementation of handlers.
func (a *App) HandleFunc(method string, path string, handler HandlerFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		v := Values{
			TraceID: uuid.New().String(),
			Now:     time.Now(),
		}

		ctx := context.WithValue(r.Context(), KeyValues, &v)
		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
			return
		}
	}
	a.Router.HandleFunc(path, h).Methods(method)

}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}