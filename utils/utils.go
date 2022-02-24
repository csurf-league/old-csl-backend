package utils

import (
	"encoding/json"
	"net/http"

	"github.com/robyzzz/csl-backend/model"
	"github.com/solovev/steam_go"
)

type ErrorResponse struct {
	Code     int
	ErrorMsg string
}

// Error response in JSON format
func APIErrorRespond(w http.ResponseWriter, res ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// Converts PlayerSummaries to SteamUser
func PlayerSummariesToSteamUser(user *steam_go.PlayerSummaries) model.SteamUser {
	return model.SteamUser{
		ID:             0,
		SteamID:        user.SteamId,
		PersonaName:    user.PersonaName,
		LastLogOff:     user.LastLogOff,
		ProfileUrl:     user.ProfileUrl,
		Avatar:         user.Avatar,
		AvatarMedium:   user.AvatarMedium,
		AvatarFull:     user.AvatarFull,
		RealName:       user.RealName,
		PrimaryClanID:  user.PrimaryClanId,
		TimeCreated:    user.TimeCreated,
		LocCountryCode: user.LocCountryCode,
		GameID:         user.GameId,
	}
}

type PlayerSumm steam_go.PlayerSummaries

// Converts PlayerSummaries to SteamUser
func (user *PlayerSumm) ToSteamUser() model.SteamUser {
	return model.SteamUser{
		ID:             0,
		SteamID:        user.SteamId,
		PersonaName:    user.PersonaName,
		LastLogOff:     user.LastLogOff,
		ProfileUrl:     user.ProfileUrl,
		Avatar:         user.Avatar,
		AvatarMedium:   user.AvatarMedium,
		AvatarFull:     user.AvatarFull,
		RealName:       user.RealName,
		PrimaryClanID:  user.PrimaryClanId,
		TimeCreated:    user.TimeCreated,
		LocCountryCode: user.LocCountryCode,
		GameID:         user.GameId,
	}
}
