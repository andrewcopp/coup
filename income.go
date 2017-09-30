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
