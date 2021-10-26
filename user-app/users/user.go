package users

import (
	"fmt"
	"net/http"
	appErr "user-app/apErr"
)

type User struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
}

var users = map[uint64]*User{

	123: &User{
		FName: "Ajay",
		LName: "Sharma",
		Email: "ajay@email.com",
	},
}

func FetchUser(id uint64) (*User, *appErr.ApplicationError) {

	u := users[id] // *User default value = nil

	if u != nil { // user is found
		return u, nil
	}
	return nil, &appErr.ApplicationError{
		Msg:        fmt.Sprintf("user id not found %v", id),
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}
}
