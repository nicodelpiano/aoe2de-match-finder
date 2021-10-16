package api

import "github.com/nicodelpiano/aoe2de-match-finder/game"

type matchpredicate func(Match) bool
type playerpredicate func(Player) bool

type composer func(matchpredicate) matchpredicate

var compose = func(predicates ...matchpredicate) matchpredicate {
	return func(m Match) bool {
		for _, predicate := range predicates {
			if !predicate(m) {
				return false
			}
		}
		return true
	}
}

func playerVictory(playerNumber int) matchpredicate {
	return func(m Match) bool {
		player, err := m.GetPlayer(playerNumber)
		if err != nil {
			return false
		}
		return player.Won
	}
}

func playerNameVictory(playerName string) matchpredicate {
	return func(m Match) bool {
		player, err := m.GetPlayerByName(playerName)
		if err != nil {
			return false
		}
		return player.Won
	}
}

func playerCiv(playerNumber int) func(c game.Civ) matchpredicate {
	return func(c game.Civ) matchpredicate {
		return func(m Match) bool {
			player, err := m.GetPlayer(playerNumber)
			if err != nil {
				return false
			}
			if !c.IsValid() {
				return false
			}
			return player.Civ == c.ID
		}
	}
}

func playerNameCiv(playerName string) func(c game.Civ) matchpredicate {
	return func(c game.Civ) matchpredicate {
		return func(m Match) bool {
			player, err := m.GetPlayerByName(playerName)
			if err != nil {
				return false
			}
			if !c.IsValid() {
				return false
			}
			return player.Civ == c.ID
		}
	}
}

func numPlayers(numPlayers int) matchpredicate {
	return func(m Match) bool {
		return m.NumPlayers == numPlayers
	}
}

var p1Victory matchpredicate = playerVictory(1)
var p2Victory matchpredicate = playerVictory(2)

var p1CivIs func(c game.Civ) matchpredicate = playerCiv(1)
var p2CivIs func(c game.Civ) matchpredicate = playerCiv(2)

var only2Players matchpredicate = numPlayers(2)
var rankedGame matchpredicate = func(m Match) bool {
	return m.Ranked
}
var any matchpredicate = func(m Match) bool {
	return true
}

var PlayerWonWithFranks = func(playerName string) matchpredicate {
	return compose(only2Players, rankedGame, playerNameVictory(playerName), playerNameCiv(playerName)(game.Franks))
}
