package controller

import (
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	steamid := config.GetSessionID(r)
	if steamid == "" {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, "Invalid session ID."))
		return
	}
	println("oi")
	w.Write([]byte(steamid))
}
