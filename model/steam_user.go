package model

import (
	"log"
)

type SteamUser struct {
	ID             uint64 `json:"id"`
	SteamID        string `json:"steamid"`
	PersonaName    string `json:"personaname"`
	LastLogOff     int    `json:"lastlogoff"`
	ProfileUrl     string `json:"profileurl"`
	Avatar         string `json:"avatar"`
	AvatarMedium   string `json:"avatarmedium"`
	AvatarFull     string `json:"avatarfull"`
	RealName       string `json:"realname"`
	PrimaryClanID  string `json:"primaryclanid"`
	TimeCreated    int    `json:"timecreated"`
	LocCountryCode string `json:"loccountrycode"`
	GameID         int    `json:"gameid"`
}

func CreateSteamUser(user SteamUser) error {
	userExists, err := DoesSteamUserExist(user.SteamID)
	if err != nil {
		return err
	}

	if userExists {
		return nil
	}

	query := `INSERT INTO steam_user(steamid, personaname, lastlogoff, profileurl, avatar, avatarmedium, avatarfull, realname, primaryclanid, timecreated, loccountrycode, gameid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

	_, err = db.Exec(query, user.SteamID, user.PersonaName, user.LastLogOff, user.ProfileUrl, user.Avatar, user.AvatarMedium, user.AvatarFull, user.RealName, user.PrimaryClanID, user.TimeCreated, user.LocCountryCode, user.GameID)

	if err != nil {
		log.Printf("Error creating steam user @CreateSteamUser: %s\n", err.Error())
		return err
	}

	return nil
}

func UpdateSteamUser(user SteamUser) error {
	query := `UPDATE steam_user SET personaname = $1::text, lastlogoff = $2, profileurl = $3::text, avatar = $4::text, avatarmedium = $5::text, avatarfull = $6::text, realname = $7::text, primaryclanid = $8::text, timecreated = $9, loccountrycode = $10::text, gameid = $11 WHERE steamid=$12::text ;`

	_, err = db.Exec(query, user.PersonaName, user.LastLogOff, user.ProfileUrl, user.Avatar, user.AvatarMedium, user.AvatarFull, user.RealName, user.PrimaryClanID, user.TimeCreated, user.LocCountryCode, user.GameID, user.SteamID)

	if err != nil {
		log.Printf("Error updating steam user @UpdateSteamUser: %s\n", err.Error())
		return err
	}

	return nil
}

func DoesSteamUserExist(steamid string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM steam_user WHERE steamid = '" + steamid + "';").Scan(&count)
	if err != nil {
		log.Printf("Error querying steam user @DoesSteamUserExist: %s\n", err.Error())
		return false, err
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}