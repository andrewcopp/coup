package coup

import "fmt"

type Income struct {
	Subject *Player
}

func NewIncome(sub *Player) *Move {
	income := Income{
		Subject: sub,
	}

	return NewMove(
		fmt.Sprintf("%s takes income.", sub.Name),
		0,
		nil,
		[]CardType{},
		income.Resolve,
	)
}

func (i *Income) Resolve(state *State) {
	i.Subject.Coins++
	Account(i.Subject)
}
