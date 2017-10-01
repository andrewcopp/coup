package coup

import "fmt"

func NewAssassinate(sub *Player, obj *Player, state *State) *Move {
	return NewMove(
		fmt.Sprintf("%s assassinates %s.", sub.Name, obj.Name),
		3,
		NewClaim(sub, Assassin, obj),
		[]CardType{Contessa},
		AssassinateFunc(sub, obj, state),
	)
}

func AssassinateFunc(sub *Player, obj *Player, state *State) func() {
	return func() {
		if len(obj.Hand) != 0 {
			revealed := obj.Reveal(state, nil)
			state.Discard.Add(revealed)
			fmt.Printf("%s discards a %s.\n", obj.Name, revealed.Name())
		}
	}
}
