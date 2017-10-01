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
		fmt.Sprintf("%s assassinates %s.", sub.Name, obj.Name),
		3,
		NewClaim(sub, Assassin, obj),
		[]Type{Contessa},
		assassinate.Resolve,
	)
}

func (a *Assassinate) Resolve(state *State) {
	if len(a.Object.Hand) != 0 {
		revealed := a.Object.Reveal(state, nil)
		state.Revealed = append(state.Revealed, revealed)
		fmt.Printf("%s discards a %s.\n", a.Object.Name, revealed.Name())
	}
}
