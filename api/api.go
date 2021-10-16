package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	aoe2netAPIURL       = `https://aoe2.net/api/`
	leaderboardEndpoint = aoe2netAPIURL + `leaderboard?game=aoe2de&leaderboard_id=3`
	matchesEndpoint     = aoe2netAPIURL + `player/matches?game=aoe2de&start=0&count=10`
	// https://aoe2.net/api/match?match_id=122104697
	matchEndpoint = aoe2netAPIURL + `/match?game=aoe2de`
)

func FetchMatches(profileID int) (Matches, error) {
	matches := Matches{}
	r, err := http.Get(fmt.Sprintf("%s%s%d", matchesEndpoint, "&profile_id=", profileID))
	if err != nil {
		return matches, err
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&matches); err != nil {
		return matches, err
	}

	return matches, nil
}

func FetchMatch(matchID int) (Match, error) {
	match := Match{}
	r, err := http.Get(fmt.Sprintf("%s%s%d", matchEndpoint, "&match_id=", matchID))
	if err != nil {
		return match, err
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		return match, err
	}

	return match, nil
}

func FetchLeaderboard(count int) (LeaderboardResponse, error) {
	leaderboard := LeaderboardResponse{}
	r, err := http.Get(fmt.Sprintf("%s%s%d", leaderboardEndpoint, `&count=`, count))
	if err != nil {
		return leaderboard, err
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&leaderboard); err != nil {
		return leaderboard, err
	}

	// p1WonWithFranks := matches.Filter(compose(only2Players, rankedGame, playerNameVictory("Tirano de Aldea"), playerNameCiv("Tirano de Aldea")(game.Franks)))
	// log.Printf("Matches count %d", len(p1WonWithFranks))
	return leaderboard, nil
}
