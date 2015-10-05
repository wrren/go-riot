package riot

type Spell struct {
	Key rune
	Description string

	cost []int
}

func (s Spell) Cost( level int ) int {
	return s.cost[level-1];
}