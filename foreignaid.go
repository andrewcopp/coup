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
		foreignAid.Announce,
		foreignAid.Pay,
		nil,
		[]Type{Duke},
		foreignAid.Resolve,
	)
}

func (f *ForeignAid) Announce() {
	fmt.Printf("%s takes foreign aid.\n", f.Subject.Name)
}

func (f *ForeignAid) Pay() {}

func (f *ForeignAid) Resolve(state *State) {
	f.Subject.Coins += 2
	Account(f.Subject)
}
