package middleware

import (
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/controller"
)

// Used to let client access only routes that need him to be logged in
func IsAuthenticated(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	next := http.HandlerFunc(h)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if config.SessionAlreadyExists(r) {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
			}
		})
}

// Used to update steam user data when acessing /login
// If user is already logged in, we update, else we redirect to login page
func BeforeLogin(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
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
