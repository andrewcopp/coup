package coup

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
	valid := []*Move{}
	self := state.Players[0]

	for _, other := range state.Alive()[1:] {
		if state.Players[0].Coins >= 7 {
			valid = append(valid, NewCoup(self, other))
		}

		if state.Players[0].Coins < 10 {
			valid = append(valid, NewIncome(self))
			valid = append(valid, NewForeignAid(self))
			valid = append(valid, NewTax(self))
			valid = append(valid, NewExchange(self))
			valid = append(valid, NewSteal(self, other))
		}

		if state.Players[0].Coins >= 3 {
			valid = append(valid, NewAssassinate(self, other))
		}
	}

	move := (*p.Brain).Decide(state, valid)
	return NewAction(move)
}

func (p *Player) Reveal(state *State, t *Type) *Card {
	if t != nil {
		for i, card := range p.Hand {
			if card.Type == *t {
				p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
				return card
			}
		}
	}

	return (*p.Brain).Discard(state)
}

func (p *Player) Dispute(claim *Claim) {
	if (*p.Brain).Dispute(claim) {
		claim.Challenge = NewChallenge(p)
	}
}
