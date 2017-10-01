package coup

import "fmt"

func NewIncome(sub *Player) *Move {
	return NewMove(
		Income,
		sub,
		fmt.Sprintf("%s takes income.", sub.Name),
		0,
		nil,
		[]CardType{},
		IncomeFunc(sub),
	)
}

func IncomeFunc(sub *Player) func() {
	return func() {
		sub.Coins++
		Account(sub)
	}
}
