package coup

import (
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	Brain    *Decider
	Coins    int
	Revealed []*Card
	Hidden   []*Card
}

func NewPlayer(name string, brain *Decider, coins int, revealed []*Card, hidden []*Card) *Player {
	return &Player{
		Name:     name,
		Brain:    brain,
		Coins:    coins,
		Revealed: revealed,
		Hidden:   hidden,
	}
}

func (p *Player) Copy() *Player {

	revealed := make([]*Card, len(p.Revealed))
	for i, card := range p.Revealed {
		revealed[i] = card.Copy()
	}

	hidden := make([]*Card, len(p.Hidden))
	for i, card := range p.Hidden {
		hidden[i] = card.Copy()
	}

	return NewPlayer(p.Name, p.Brain, p.Coins, revealed, hidden)

}

func (p *Player) Move(state *State) *Action {
	return (*p.Brain).Decide(state)
}

func (p *Player) Produce(t Type) *Card {
	for i, card := range p.Hidden {
		if card.Type == t {
			p.Hidden = append(p.Hidden[:i], p.Hidden[i+1:]...)
			return card
		}
	}

	return nil
}

func (p *Player) Discard() *Card {
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(p.Hidden))
	card := p.Hidden[i]
	p.Hidden = append(p.Hidden[:i], p.Hidden[i+1:]...)
	return card
}

func (p *Player) Reveal() {
	p.Revealed = append(p.Revealed, p.Discard())
}
