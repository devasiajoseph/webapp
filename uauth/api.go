package uauth

import (
	"errors"
	"net/http"
)

func GetAuthenticatedUser(r *http.Request) (AuthUser, error) {
	var authUser AuthUser
	cookie, err := r.Cookie("uauth-token")
	if err != nil {
		return authUser, errors.New("Unauthorized")
	}
	if cookie == nil {
		return authUser, errors.New("Unauthorized")
	}
	authUser, err = GetUserByAuth(cookie.Value)
	if err != nil {
		return authUser, err
	}
	go keepAlive(cookie.Value)
	return authUser, err
}
