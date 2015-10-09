package riot

import (
	"time"
)

type ChampionStatus struct {
	retrieved time.Time
	active bool
	botEnabled bool
	botMMEnabled bool
	freeToPlay bool
	id uint64
	rankedPlayEnabled bool
}

type Champion struct {
	id uint64		`json:"id"`
	name string		`json:"name"`
	title string		`json:"title"`
	status *ChampionStatus
}

// Champion API Endpoint
const ChampionEndpoint = "champion"

// Version for lol-static champion data retrieval
const ChampionStaticVersion	= "1.2"
// Version for champion realtime status endpoint
const ChampionStatusVersion	= "1.2"

func NewChampion( id uint64, name string, title string ) *Champion {
	return &Champion{ id, name, title, nil }
}

//
//	Get a list of all champions in the game
//
func (r Riot) GetChampions() ( error, []Champion ) {
	var json interface{}
	champions := make([]Champion, 1)
	
	err := r.GetAndUnmarshal(true, ChampionEndpoint, ChampionStaticVersion, make([]string, 0, 0), make(map[string]string), &json)

	if err != nil {
		return err, champions
	}

	m := json.(map[string]interface{})

	for name, d := range m {
		data := d.(map[string]interface{})

		champions = append(champions, *NewChampion(data["id"].(uint64), data["name"].(string), data["title"].(string)))
	}

	return nil, champions
}

//
//	Get the champion with the specified ID
//
func (r Riot) GetChampionByID( id int ) ( error, Champion ) {

}

//
//	Get the champion with the specified name
//
func (r Riot) GetChampionByName( name string ) ( error, Champion ) {

}

//
//	Get the current status of the specified champion
//
func ( c *Champion ) Status( maxAge time.Duration ) {

}