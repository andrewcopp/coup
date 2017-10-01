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
		3,
		NewClaim(sub, Assassin, obj),
		[]Type{Contessa},
		assassinate.Resolve,
	)
}

func (a *Assassinate) Announce() {
	fmt.Printf("%s assassinates %s.\n", a.Subject.Name, a.Object.Name)
}

func (a *Assassinate) Resolve(state *State) {
	if len(a.Object.Hand) != 0 {
		a.Object.Reveal(state, nil)
		fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
	}
}
