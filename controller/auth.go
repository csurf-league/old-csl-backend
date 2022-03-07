package controller

import (
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/utils"
	"github.com/solovev/steam_go"
)

// GET /login - redirect to steam auth and validate user
func Login(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(user)
		http.Redirect(w, r, config.FRONTEND_URL, http.StatusTemporaryRedirect)
	}
}

// GET /logout - Log out from current session
func Logout(w http.ResponseWriter, r *http.Request) {
	config.RemoveSessionID(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
