package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Matches []Match

type Match struct {
	MatchID           string      `json:"match_id"`
	LobbyID           interface{} `json:"lobby_id"`
	MatchUUID         string      `json:"match_uuid"`
	Version           string      `json:"version"`
	Name              string      `json:"name"`
	NumPlayers        int         `json:"num_players"`
	NumSlots          int         `json:"num_slots"`
	AverageRating     interface{} `json:"average_rating"`
	Cheats            bool        `json:"cheats"`
	FullTechTree      bool        `json:"full_tech_tree"`
	EndingAge         int         `json:"ending_age"`
	Expansion         interface{} `json:"expansion"`
	GameType          int         `json:"game_type"`
	HasCustomContent  interface{} `json:"has_custom_content"`
	HasPassword       bool        `json:"has_password"`
	LockSpeed         bool        `json:"lock_speed"`
	LockTeams         bool        `json:"lock_teams"`
	MapSize           int         `json:"map_size"`
	MapType           int         `json:"map_type"`
	Pop               int         `json:"pop"`
	Ranked            bool        `json:"ranked"`
	LeaderboardID     int         `json:"leaderboard_id"`
	RatingType        int         `json:"rating_type"`
	Resources         int         `json:"resources"`
	Rms               interface{} `json:"rms"`
	Scenario          interface{} `json:"scenario"`
	Server            string      `json:"server"`
	SharedExploration bool        `json:"shared_exploration"`
	Speed             int         `json:"speed"`
	StartingAge       int         `json:"starting_age"`
	TeamTogether      bool        `json:"team_together"`
	TeamPositions     bool        `json:"team_positions"`
	TreatyLength      int         `json:"treaty_length"`
	Turbo             bool        `json:"turbo"`
	Victory           int         `json:"victory"`
	VictoryTime       int         `json:"victory_time"`
	Visibility        int         `json:"visibility"`
	Opened            int         `json:"opened"`
	Started           int         `json:"started"`
	Finished          int         `json:"finished"`
	Players           Players     `json:"players"`
}

func (m *Matches) Filter(predicate func(Match) bool) Matches {
	matches := Matches{}
	for _, match := range *m {
		if predicate(match) {
			matches = append(matches, match)
		}
	}
	return matches
}

func (m *Matches) Find(predicate func(Match) bool) *Match {
	for _, match := range *m {
		if predicate(match) {
			return &match
		}
	}
	return nil
}

func (m *Match) validPlayer(playerNumber int) bool {
	return playerNumber > 0 && playerNumber <= m.NumPlayers
}

var errPlayerNotFound = fmt.Errorf("player not found")

func (m *Match) GetPlayer(playerNumber int) (*Player, error) {
	if !m.validPlayer(playerNumber) {
		return nil, errPlayerNotFound
	}
	return &m.Players[playerNumber-1], nil
}

func (m *Match) GetPlayerByName(playerName string) (*Player, error) {
	for _, player := range m.Players {
		if player.Name == playerName {
			return &player, nil
		}
	}
	return nil, errPlayerNotFound
}

func (m *Match) Info() string {
	p1, _ := m.GetPlayer(1)
	p2, _ := m.GetPlayer(2)
	return fmt.Sprintf("MatchID: %s\nMap: %d\nPlayer: %s\nOponent: %s\n", m.MatchID, m.MapType, p1.Name, p2.Name)
}

const (
	gameReplayEndpoint = "https://aoe.ms/replay/"
)

func (m *Match) Download() error {
	p1, _ := m.GetPlayer(1)

	// https://aoe.ms/replay/?gameId=122104697&profileId=3392644
	r, err := http.Get(fmt.Sprintf("%s?gameId=%s&profileId=%d", gameReplayEndpoint, m.MatchID, p1.ProfileID))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Create the file
	out, err := os.Create(fmt.Sprintf("replay-%s-%d", m.MatchID, p1.ProfileID))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, r.Body)
	return err
}
