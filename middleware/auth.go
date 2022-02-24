package middleware

import (
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/controller"
)

// Used to update steam user data when acessing a route or to let them log in
// If user is already logged in, we update, else we redirect to login page
// Usage: /login
func IsLogged(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	next := http.HandlerFunc(h)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if config.SessionAlreadyExists(r) {
				controller.UpdateSteamUser(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
}
