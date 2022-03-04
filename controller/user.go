package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/utils"
)

// GET /api/users/ - returns all users data
func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := model.UserIndex()
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GET /api/user/{steamid} - returns player's data by steamid
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := model.GetUser(mux.Vars(r)["steamid"])
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GET /profile - returns user stats if authenticated
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := model.GetUser(config.GetSessionID(r))
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}
