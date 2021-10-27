package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"service-app/web"
)

func (m *Mid) Panics(next web.HandlerFunc) web.HandlerFunc {

	// This is the actual middleware function to be executed.
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

		// If the context is missing this value, request the service
		// to be shutdown gracefully.
		v, ok := ctx.Value(web.KeyValues).(*web.Values)
		if !ok {
			return web.NewShutdownError("web value missing from context")
		}


		// Defer a function to recover from a panic.
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(fmt.Sprintf("PANIC: %v", r))

				// Log the Go stack trace for this panic'd goroutine.
				log.Printf("%s :\n%s", v.TraceID, debug.Stack())
			}
		}()

		// Call the next HandlerFunc and set its return value in the err variable.
		return next(ctx, w, r)
	}

}