package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	//return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	//return errors.New("not trusted")
	//return web.NewShutdownError("i want to shutdown")
	//panic("i want to panic")
	status := struct {
		Status string
	}{
		Status: "ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
