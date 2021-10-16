package api

type Players []Player

type Player struct {
	ProfileID    int         `json:"profile_id"`
	SteamID      string      `json:"steam_id"`
	Name         string      `json:"name"`
	Clan         interface{} `json:"clan"`
	Country      string      `json:"country"`
	Slot         int         `json:"slot"`
	SlotType     int         `json:"slot_type"`
	Rating       int         `json:"rating"`
	RatingChange interface{} `json:"rating_change"`
	Games        interface{} `json:"games"`
	Wins         interface{} `json:"wins"`
	Streak       interface{} `json:"streak"`
	Drops        interface{} `json:"drops"`
	Color        int         `json:"color"`
	Team         int         `json:"team"`
	Civ          int         `json:"civ"`
	CivAlpha     int         `json:"civ_alpha"`
	Won          bool        `json:"won"`
}
