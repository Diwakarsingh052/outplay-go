package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	appErr "user-app/appErr"
	"user-app/users"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		appEr := appErr.ApplicationError{
			Msg:        "user Id must be a number",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}
		jsonData, err := json.Marshal(appEr)
		if err != nil {
			http.Error(w, "request not processed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(appEr.StatusCode)
		w.Write(jsonData)
		// json.NewEncoder.Encode(struct)
		return

	}
	user, appErr := users.FetchUser(userId)

	if appErr != nil {
		jsonErrData, err := json.Marshal(appErr)

		if err != nil {
			log.Println(err)
			http.Error(w, "Request cannot be processed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonErrData)
		return
	}

	jsonData, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return

}
