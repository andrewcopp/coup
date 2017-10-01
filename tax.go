package coup

import "fmt"

type Tax struct {
	Subject   *Player
	Challenge *Challenge
}

func NewTax(sub *Player) *Move {
	tax := Tax{
		Subject:   sub,
		Challenge: nil,
	}

	return NewMove(
		fmt.Sprintf("%s taxes.", sub.Name),
		0,
		NewClaim(sub, Duke, nil),
		[]Type{},
		tax.Resolve,
	)
}

func (t *Tax) Resolve(state *State) {
	t.Subject.Coins += 3
	Account(t.Subject)
}
