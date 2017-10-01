package coup

import "fmt"

func NewTax(sub *Player) *Move {
	return NewMove(
		Tax,
		sub,
		fmt.Sprintf("%s taxes.", sub.Name),
		0,
		NewClaim(sub, Duke, nil),
		[]CardType{},
		TaxFunc(sub),
	)
}

func TaxFunc(sub *Player) func() {
	return func() {
		sub.Coins += 3
		Account(sub)
	}
}
