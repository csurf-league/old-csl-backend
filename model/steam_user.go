package model

import (
	"log"
)

// SteamUser represents steam user's relevant data
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
	Created_At     string `json:"created_at"`
	Updated_At     string `json:"updated_at"`
}

// Returns a user from database by steamid
func GetSteamUser(steamid string) (SteamUser, error) {
	user := SteamUser{}
	err := db.Get(&user, "SELECT * FROM steam_user WHERE steamid = $1;", steamid)
	return user, err
}

// Creates a new user on the database
func CreateSteamUser(user SteamUser) error {
	userExists, err := DoesSteamUserExist(user.SteamID)
	if err != nil {
		return err
	}

	if userExists {
		return UpdateSteamUser(user)
	}

	query := `INSERT INTO steam_user(steamid, personaname, lastlogoff, profileurl, avatar, avatarmedium, 
		avatarfull, realname, primaryclanid, timecreated, loccountrycode, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

	_, err = db.Exec(query, user.SteamID, user.PersonaName, user.LastLogOff, user.ProfileUrl, user.Avatar, user.AvatarMedium,
		user.AvatarFull, user.RealName, user.PrimaryClanID, user.TimeCreated, user.LocCountryCode, "NOW()")

	if err != nil {
		log.Printf("Error creating SteamUser @CreateSteamUser: %s\n", err.Error())
		return err
	}

	err = CreatePlayer(user.SteamID)
	if err != nil {
		log.Printf("Error creating Player @CreateSteamUser: %s\n", err.Error())
		return err
	}

	return nil
}

// Updates user from the database
func UpdateSteamUser(user SteamUser) error {
	query := `UPDATE steam_user SET personaname = $1::text, lastlogoff = $2, profileurl = $3::text, 
		avatar = $4::text, avatarmedium = $5::text, avatarfull = $6::text, realname = $7::text, 
		primaryclanid = $8::text, timecreated = $9, loccountrycode = $10::text, updated_at = NOW() 
		WHERE steamid=$11::text ;`

	_, err = db.Exec(query, user.PersonaName, user.LastLogOff, user.ProfileUrl,
		user.Avatar, user.AvatarMedium, user.AvatarFull, user.RealName,
		user.PrimaryClanID, user.TimeCreated, user.LocCountryCode, user.SteamID)

	if err != nil {
		log.Printf("Error updating SteamUser @UpdateSteamUser: %s\n", err.Error())
		return err
	}

	return nil
}

// Returns true if user by steamid already exists
func DoesSteamUserExist(steamid string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM steam_user WHERE steamid = '" + steamid + "';").Scan(&count)
	if err != nil {
		log.Printf("Error querying SteamUser @DoesSteamUserExist: %s\n", err.Error())
		return false, err
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
