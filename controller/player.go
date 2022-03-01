package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/utils"
)

// GET /api/playerstats/{steamid} - returns {steamid}'s stats or 404 if not found
func GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	player, err := model.GetPlayerStats(mux.Vars(r)["steamid"])
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(player)
}
