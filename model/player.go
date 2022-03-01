package model

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

func GetPlayerStats(steamID string) (PlayerStats, error) {
	user := PlayerStats{}
	err := db.Get(&user, "SELECT * FROM player_stats WHERE player_steamid = $1;", steamID)
	return user, err
}
