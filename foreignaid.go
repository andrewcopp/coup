package coup

import "fmt"

func NewForeignAid(sub *Player) *Move {
	return NewMove(
		fmt.Sprintf("%s takes foreign aid.", sub.Name),
		0,
		nil,
		[]CardType{Duke},
		ForeignAidFunc(sub),
	)
}

func ForeignAidFunc(sub *Player) func() {
	return func() {
		sub.Coins += 2
		Account(sub)
	}
}
