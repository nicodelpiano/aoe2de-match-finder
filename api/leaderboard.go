package api

import "fmt"

type LeaderboardResponse struct {
	Total         int         `json:"total"`
	LeaderboardID int         `json:"leaderboard_id"`
	Start         int         `json:"start"`
	Count         int         `json:"count"`
	Leaderboard   Leaderboard `json:"leaderboard"`
}

type Leaderboard []LeaderboardEntry

type LeaderboardEntry struct {
	ProfileID      int         `json:"profile_id"`
	Rank           int         `json:"rank"`
	Rating         int         `json:"rating"`
	SteamID        string      `json:"steam_id"`
	Icon           interface{} `json:"icon"`
	Name           string      `json:"name"`
	Clan           interface{} `json:"clan"`
	Country        string      `json:"country"`
	PreviousRating int         `json:"previous_rating"`
	HighestRating  int         `json:"highest_rating"`
	Streak         int         `json:"streak"`
	LowestStreak   int         `json:"lowest_streak"`
	HighestStreak  int         `json:"highest_streak"`
	Games          int         `json:"games"`
	Wins           int         `json:"wins"`
	Losses         int         `json:"losses"`
	Drops          int         `json:"drops"`
	LastMatch      int         `json:"last_match"`
	LastMatchTime  int         `json:"last_match_time"`
}

func (le *LeaderboardEntry) Info() string {
	return fmt.Sprintf(`ProfileID: %d
Rank: %d
Rating: %d
Name: %s`, le.ProfileID, le.Rank, le.Rating, le.Name)
}
