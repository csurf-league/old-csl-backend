package controller

import (
	"log"

	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/utils"
	"github.com/solovev/steam_go"
)

func CreateSteamUser(user *steam_go.PlayerSummaries) error {
	return model.CreateSteamUser(utils.PlayerSummariesToSteamUser(user))
}

func UpdateSteamUser(steamID string) error {
	updatedUser, err := steam_go.GetPlayerSummaries(steamID, config.STEAM_API_KEY)
	if err != nil {
		log.Printf("Failed to get PlayerSummaries @UpdateSteamUser: %s", err.Error())
		return err
	}

	return model.UpdateSteamUser(utils.PlayerSummariesToSteamUser(updatedUser))
}
