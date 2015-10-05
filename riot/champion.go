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
	name string
	title string
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
func GetChampions() ( error, []Champion ) {

}

//
//	Get the champion with the specified ID
//
func GetChampionByID( id int ) ( error, Champion ) {

}

//
//	Get the champion with the specified name
//
func GetChampionByName( name string ) ( error, Champion ) {

}

//
//	Get the current status of the specified champion
//
func ( c *Champion ) Status( maxAge time.Duration ) {

}