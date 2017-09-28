package coup

type Player struct {
	Name     string
	Brain    *Decider
	Coins    int
	Revealed []Card
	Hidden   []Card
}

func NewPlayer(name string, brain *Decider, coins int, revealed []Card, hidden []Card) *Player {
	return &Player{
		Name:     name,
		Brain:    brain,
		Coins:    coins,
		Revealed: revealed,
		Hidden:   hidden,
	}
}

func (p *Player) Copy() *Player {

	revealed := make([]Card, len(p.Revealed))
	_ = copy(revealed, p.Revealed)

	hidden := make([]Card, len(p.Hidden))
	_ = copy(hidden, p.Hidden)

	return NewPlayer(p.Name, p.Brain, p.Coins, revealed, hidden)

}

func (p *Player) Move(state *State) *Action {
	return (*p.Brain).Decide(state)
}
