package coup

import "fmt"

type Income struct {
	Subject *Player
}

func NewIncome(sub *Player) *Income {
	return &Income{
		Subject: sub,
	}
}

func (i *Income) Modify(state *State) {
	i.Subject.Coins++
	fmt.Printf("%s takes income.\n", i.Subject.Name)
	Account(i.Subject)
}
