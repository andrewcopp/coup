package coup

import (
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Brain *Decider
	Coins int
	Hand  []*Card
}

func NewPlayer(name string, brain *Decider, coins int, hand []*Card) *Player {
	return &Player{
		Name:  name,
		Brain: brain,
		Coins: coins,
		Hand:  hand,
	}
}

func (p *Player) Copy() *Player {

	hand := make([]*Card, len(p.Hand))
	for i, card := range p.Hand {
		hand[i] = card.Copy()
	}

	return NewPlayer(p.Name, p.Brain, p.Coins, hand)

}

func (p *Player) Move(state *State) *Action {
	return (*p.Brain).Decide(state)
}

func (p *Player) Produce(t Type) *Card {
	for i, card := range p.Hand {
		if card.Type == t {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			return card
		}
	}

	return nil
}

func (p *Player) Discard() *Card {
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(p.Hand))
	card := p.Hand[i]
	p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
	return card
}

func (p *Player) Reveal(state *State) {
	state.Revealed = append(state.Revealed, p.Discard())
}

type Claim struct {
	Subject    *Player
	Declared   Type
	Object     *Player
	Challenger *Player
	Revealed   *Type
}

func NewClaim(sub *Player, dec Type, obj *Player) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
		Object:   obj,
	}
}

func (p *Player) Claim(t Type) {

}

func (p *Player) Dispute(claim *Claim) {
	if (*p.Brain).Dispute(claim) {
		claim.Challenger = p
	}
}
