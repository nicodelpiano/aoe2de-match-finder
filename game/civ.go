package game

type Civ struct {
	ID     int    `json:"id"`
	String string `json:"string"`
}

var Civs = map[int]Civ{
	1:  Britons,
	2:  Franks,
	3:  Goths,
	4:  Teutons,
	5:  Japanese,
	6:  Chinese,
	7:  Byzantines,
	8:  Persians,
	9:  Saracens,
	10: Turks,
	11: Vikings,
	12: Mongols,
	13: Celts,
	14: Spanish,
	15: Aztecs,
	16: Mayans,
	17: Huns,
	18: Koreans,
	19: Italians,
	20: Indians,
	21: Incas,
	22: Magyars,
	23: Slavs,
	24: Portuguese,
	25: Ethiopians,
	26: Malians,
	27: Berbers,
	28: Khmer,
	29: Malay,
	30: Burmese,
	31: Vietnamese,
	32: Bulgarians,
	33: Tatars,
	34: Cumans,
	35: Lithuanians,
	36: Burgundians,
	37: Sicilians,
	38: Poles,
	39: Bohemians,
}

var Britons = Civ{ID: 1, String: "Britons"}
var Franks = Civ{ID: 2, String: "Franks"}
var Goths = Civ{ID: 3, String: "Goths"}
var Teutons = Civ{ID: 4, String: "Teutons"}
var Japanese = Civ{ID: 5, String: "Japanese"}
var Chinese = Civ{ID: 6, String: "Chinese"}
var Byzantines = Civ{ID: 7, String: "Byzantines"}
var Persians = Civ{ID: 8, String: "Persians"}
var Saracens = Civ{ID: 9, String: "Saracens"}
var Turks = Civ{ID: 10, String: "Turks"}
var Vikings = Civ{ID: 11, String: "Vikings"}
var Mongols = Civ{ID: 12, String: "Mongols"}
var Celts = Civ{ID: 13, String: "Celts"}
var Spanish = Civ{ID: 14, String: "Spanish"}
var Aztecs = Civ{ID: 15, String: "Aztecs"}
var Mayans = Civ{ID: 16, String: "Mayans"}
var Huns = Civ{ID: 17, String: "Huns"}
var Koreans = Civ{ID: 18, String: "Koreans"}
var Italians = Civ{ID: 19, String: "Italians"}
var Indians = Civ{ID: 20, String: "Indians"}
var Incas = Civ{ID: 21, String: "Incas"}
var Magyars = Civ{ID: 22, String: "Magyars"}
var Slavs = Civ{ID: 23, String: "Slavs"}
var Portuguese = Civ{ID: 24, String: "Portuguese"}
var Ethiopians = Civ{ID: 25, String: "Ethiopians"}
var Malians = Civ{ID: 26, String: "Malians"}
var Berbers = Civ{ID: 27, String: "Berbers"}
var Khmer = Civ{ID: 28, String: "Khmer"}
var Malay = Civ{ID: 29, String: "Malay"}
var Burmese = Civ{ID: 30, String: "Burmese"}
var Vietnamese = Civ{ID: 31, String: "Vietnamese"}
var Bulgarians = Civ{ID: 32, String: "Bulgarians"}
var Tatars = Civ{ID: 33, String: "Tatars"}
var Cumans = Civ{ID: 34, String: "Cumans"}
var Lithuanians = Civ{ID: 35, String: "Lithuanians"}
var Burgundians = Civ{ID: 36, String: "Burgundians"}
var Sicilians = Civ{ID: 37, String: "Sicilians"}
var Poles = Civ{ID: 38, String: "Poles"}
var Bohemians = Civ{ID: 39, String: "Bohemians"}

func (c Civ) IsValid() bool {
	for _, civ := range Civs {
		if c.String == civ.String && c.ID == civ.ID {
			return true
		}
	}
	return false
}
