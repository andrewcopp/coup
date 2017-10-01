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
		tax.Announce,
		tax.Pay,
		NewClaim(sub, Duke, nil),
		[]Type{},
		tax.Resolve,
	)
}

func (t *Tax) Announce() {
	fmt.Printf("%s taxes.\n", t.Subject.Name)
}

func (t *Tax) Pay() {}

func (t *Tax) Resolve(state *State) {
	t.Subject.Coins += 3
	Account(t.Subject)
}
