package model

// PlayerStats represents game data
type PlayerStats struct {
	ID             uint64 `json:"id"`
	Player_SteamID string `json:"player_steamid"`
	Map_ID         int    `json:"map_id"`
	Kills          int    `json:"kills"`
	Deaths         int    `json:"deaths"`
	Assists        int    `json:"assists"`
	Shots          int    `json:"shots"`
	Hits           int    `json:"hits"`
	Damage         int    `json:"damage"`
	First_Blood    int    `json:"first_blood"`
	Aces           int    `json:"aces"`
	Headshots      int    `json:"headshots"`
	No_Scope       int    `json:"no_scope"`
	Count          int    `json:"count"`
	Playtime       int    `json:"playtime"`
	Match_Win      int    `json:"match_win"`
	Match_Lose     int    `json:"match_lose"`
	Match_Draw     int    `json:"match_draw"`
	Rounds_Won     int    `json:"rounds_won"`
	Rounds_Lost    int    `json:"rounds_lost"`
	Mvp            int    `json:"mvp"`
}

// Get all players.
func GetAllPlayers() ([]PlayerStats, error) {
	players := []PlayerStats{}
	err := db.Select(&players, "SELECT * FROM player_stats;")
	return players, err
}

// Find player by steamid.
func GetPlayer(steamid string) (PlayerStats, error) {
	player := PlayerStats{}
	err := db.Get(&player, "SELECT * FROM player_stats WHERE player_steamid = $1;", steamid)
	return player, err
}

// Creates 1 record on player_stats.
// Useful for creating the user when he logs in for the first time.
func CreateDefaultPlayer(steamid string) error {
	query := `INSERT INTO player_stats(player_steamid, map_id) VALUES ($1,1);`
	_, err = db.Exec(query, steamid)
	return err
}
