package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/utils"
	"github.com/solovev/steam_go"
)

// GET /login - redirect to steam auth and validate user
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", config.FRONTEND_URL)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	
	opId := steam_go.NewOpenId(r)

	switch opId.Mode() {
	case "cancel":
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	case "":
		// redirect to steam auth
		http.Redirect(w, r, opId.AuthUrl(), http.StatusTemporaryRedirect)
	default:
		// login success
		w.Header().Set("Content-Type", "application/json")

		user, err := opId.ValidateAndGetUser(config.STEAM_API_KEY)
		if err != nil {
			utils.APIErrorRespond(w, utils.NewAPIError(http.StatusInternalServerError, "ValidateAndGetUser: "+err.Error()))
			return
		}

		if err = CreateSteamUser(user); err != nil {
			utils.APIErrorRespond(w, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
			return
		}
		
		config.CreateSessionID(w, r, user.SteamId)
		http.Redirect(w, r, config.FRONTEND_URL, http.StatusTemporaryRedirect)
	}
}

// GET /logout - Log out from current session
func Logout(w http.ResponseWriter, r *http.Request) {
	config.RemoveSessionID(w, r)

	c := &http.Cookie{
		Name:     config.SESSION_NAME,
		Value:    "",
		Path:     "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(w, c)
	
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// PUT /auth
func Auth(w http.ResponseWriter, r *http.Request) {
	id:= config.GetSessionID(r)

	user, err := model.GetSteamUser(id)
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusUnauthorized, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}