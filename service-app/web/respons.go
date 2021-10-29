package web

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	v, ok := ctx.Value(KeyValues).(*Values)
	if !ok {
		return NewShutdownError("web value missing from context")
	}

	v.StatusCode = statusCode

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// RespondError sends an error reponse back to the client.
//this should not be directly called by handler func. handler func should return errors and let middleware handle that
func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {
	var webErr *Error
	// this happens when I have my trusted error
	if ok := errors.As(err, &webErr); ok {
		er := ErrorResponse{
			Error: webErr.Err.Error(),
		}
		if err := Respond(ctx, w, er, webErr.Status); err != nil {

			return err

		}
		return nil

	}

	er := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}
	if err := Respond(ctx, w, er, http.StatusInternalServerError); err != nil {
		return err
	}
	return nil

}

func Decode(r *http.Request, val interface{}) error {
	return json.NewDecoder(r.Body).Decode(&val)
}
