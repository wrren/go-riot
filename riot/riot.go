package riot

type Region string

const (
	EUWest      Region = "euw"
	EUNorthEast        = "eune"
	NA                 = "na"
	Brazil             = "br"
	Korea              = "kr"
	LANorth            = "lan"
	LASouth            = "las"
	Oceania            = "oce"
	Turkey             = "tr"
	Russia             = "ru"
	PBE                = "pbe"
	Global             = "global"
)

type Riot struct {
	key    string
	region Region
}

func NewRiot(key string, region Region) *Riot {
	return &Riot{key, region}
}