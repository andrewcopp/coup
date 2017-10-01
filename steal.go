package coup

import "fmt"

type Steal struct {
	Subject *Player
	Object  *Player
}

func NewSteal(sub *Player, obj *Player) *Move {
	steal := Steal{
		Subject: sub,
		Object:  obj,
	}

	return NewMove(
		fmt.Sprintf("%s steals from %s.", sub.Name, obj.Name),
		0,
		NewClaim(sub, Captain, obj),
		[]CardType{Ambassador, Captain},
		steal.Resolve,
	)
}

func (s *Steal) Resolve(state *State) {
	amt := 2
	if s.Object.Coins < 2 {
		amt = s.Object.Coins
	}

	s.Object.Coins -= amt
	s.Subject.Coins += amt

	Account(s.Subject)
	Account(s.Object)
}
