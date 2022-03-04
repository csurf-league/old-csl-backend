package model

// User represents all data(player_stats + steam_user)
// TODO:
type User struct {
	ID          uint64 `json:"id"`
	SteamID     string `json:"steamid"`
	Kills       int    `json:"kills"`
	Deaths      int    `json:"deaths"`
	Assists     int    `json:"assists"`
	Shots       int    `json:"shots"`
	Hits        int    `json:"hits"`
	Damage      int    `json:"damage"`
	First_Blood int    `json:"first_blood"`
	Aces        int    `json:"aces"`
	Headshots   int    `json:"headshots"`
	No_Scope    int    `json:"no_scope"`
	Count       int    `json:"count"`
	Playtime    int    `json:"playtime"`
	Match_Win   int    `json:"match_win"`
	Match_Lose  int    `json:"match_lose"`
	Match_Draw  int    `json:"match_draw"`
	Rounds_Won  int    `json:"rounds_won"`
	Rounds_Lost int    `json:"rounds_lost"`
	Mvp         int    `json:"mvp"`
}

// Returns all players stats from db
func UserIndex() ([]User, error) {
	// TODO: get player_stats joined with steam_user stuff
	return []User{}, nil
}

// Returns player's data from db by steamid
func GetUser(steamid string) (User, error) {
	// TODO: get player_stats joined with steam_user stuff
	return User{}, err
}
