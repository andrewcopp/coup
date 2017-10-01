package coup

import "fmt"

type ForeignAid struct {
	Subject *Player
}

func NewForeignAid(sub *Player) *Move {
	foreignAid := ForeignAid{
		Subject: sub,
	}

	return NewMove(
		fmt.Sprintf("%s takes foreign aid.", sub.Name),
		0,
		nil,
		[]CardType{Duke},
		foreignAid.Resolve,
	)
}

func (f *ForeignAid) Resolve(state *State) {
	f.Subject.Coins += 2
	Account(f.Subject)
}
