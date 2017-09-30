package coup

import "fmt"

type Assassinate struct {
	Subject *Player
	Object  *Player
}

func NewAssassinate(sub *Player, obj *Player) *Move {

	assassinate := Assassinate{
		Subject: sub,
		Object:  obj,
	}

	return NewMove(
		assassinate.Announce,
		assassinate.Pay,
		NewClaim(sub, Assassin, obj),
		assassinate.Resolve,
	)
}

func (a *Assassinate) Announce() {
	fmt.Printf("%s assassinates %s.\n", a.Subject.Name, a.Object.Name)
}

func (a *Assassinate) Pay() {
	a.Subject.Coins -= 3
	Account(a.Subject)
}

func (a *Assassinate) Resolve(state *State) {
	if len(a.Object.Hand) != 0 {
		a.Object.Reveal(state)
		fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
	}
}
