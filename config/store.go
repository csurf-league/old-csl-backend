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

// Removes the auth token.
func RemoveSessionID(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, SESSION_NAME)
	session.Options.MaxAge = -1
	session.Save(r, w)
}

// Returns the auth token.
// IMPORTANT: Check if session is null (decrypting may fail so that means invalid).
func GetSessionID(r *http.Request) string {
	session, err := Store.Get(r, SESSION_NAME)
	if err != nil || session.Values["session-id"] == nil {
		return ""
	}
	return fmt.Sprintf("%s", session.Values["session-id"])
}

// Returns true if the cookie session is set.
func SessionAlreadyExists(r *http.Request) bool {
	session, _ := Store.Get(r, SESSION_NAME)
	return !session.IsNew
}
