package coup

import "fmt"

type Income struct {
	Subject *Player
}

func NewIncome(sub *Player) *Action {
	income := Income{
		Subject: sub,
	}

	return NewAction(
		income.Announce,
		income.Pay,
		income.Claim,
		income.Modify,
	)
}

func (i *Income) Announce() {
	fmt.Printf("%s takes income.\n", i.Subject.Name)
}

func (i *Income) Pay() {}

func (i *Income) Claim(state *State) bool {
	return true
}

func (i *Income) Resolve(state *State) {
	i.Subject.Coins++
	Account(i.Subject)
}

func (i *Income) Modify(state *State) {

}
