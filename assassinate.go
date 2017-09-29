package coup

import "fmt"

type Assassinate struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewAssassinate(sub *Player, obj *Player) *Assassinate {
	return &Assassinate{
		Subject:   sub,
		Object:    obj,
		Challenge: nil,
		Block:     nil,
	}
}

func (a *Assassinate) Modify(state *State) {
	a.Subject.Coins -= 3
	a.Object.Reveal()
}

func (a *Assassinate) Dispute() {

}

func (a *Assassinate) Impede() {

}

func (a *Assassinate) Describe() {
	fmt.Printf("%s assassinates %s.\n", a.Subject.Name, a.Object.Name)
	Account(a.Subject)
	fmt.Printf("%s reveals a %s.\n", a.Object.Name, a.Object.Revealed[len(a.Object.Revealed)-1].Name())
}
