package config

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("test"))

func SessionAlreadyExists(r *http.Request) bool {
	session, _ := Store.Get(r, SESSION_NAME)
	return !session.IsNew
}

func CreateSessionID(w http.ResponseWriter, r *http.Request, value string) error {
	session, _ := Store.Get(r, SESSION_NAME)
	session.Values["session-id"] = value
	return session.Save(r, w)
}
