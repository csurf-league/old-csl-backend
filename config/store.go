package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(SESSION_SECRET_KEY))

// Creates the auth token
func CreateSessionID(w http.ResponseWriter, r *http.Request, value string) error {
	session, _ := Store.Get(r, SESSION_NAME)
	session.Values["session-id"] = value
	return session.Save(r, w)
}

// Removes the auth token
func RemoveSessionID(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   SESSION_NAME,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// Returns the auth token
func GetSessionID(r *http.Request) string {
	session, _ := Store.Get(r, SESSION_NAME)
	return fmt.Sprintf("%s", session.Values["session-id"])
}

// Returns true if the cookie session is set
func SessionAlreadyExists(r *http.Request) bool {
	session, _ := Store.Get(r, SESSION_NAME)
	return !session.IsNew
}
