package utils

import (
	"github.com/robyzzz/csl-backend/model"
	"github.com/solovev/steam_go"
)

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
