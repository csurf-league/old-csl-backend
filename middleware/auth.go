package middleware

import (
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/controller"
)

// Authentication middleware called on routes that need to know if user is logged in.
// Returns to root page if session already exists.
// i.e usage: /login
func IsAuthenticated(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	next := http.HandlerFunc(h)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.SessionAlreadyExists(r) {
			err := controller.UpdateSteamUser(config.GetSessionID(r))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
