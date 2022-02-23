package store

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/robyzzz/csl-backend/utils/env"
)

var Store = sessions.NewCookieStore([]byte("test"))

func SessionAlreadyExists(r *http.Request) bool {
	session, _ := Store.Get(r, env.SESSION_NAME)
	return !session.IsNew
}

func CreateSessionID(w http.ResponseWriter, r *http.Request, value string) error {
	session, _ := Store.Get(r, env.SESSION_NAME)
	session.Values["session-id"] = value
	return session.Save(r, w)
}
