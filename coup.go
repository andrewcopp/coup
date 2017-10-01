package coup

import "fmt"

func NewCoup(sub *Player, obj *Player, state *State) *Move {
	return NewMove(
		fmt.Sprintf("%s coups %s.", sub.Name, obj.Name),
		7,
		nil,
		[]CardType{},
		CoupFunc(sub, obj, state),
	)
}

func CoupFunc(sub *Player, obj *Player, state *State) func() {
	return func() {
		revealed := obj.Reveal(state, nil)
		state.Discard.Add(revealed)
		fmt.Printf("%s reveals a %s.\n", obj.Name, revealed.Name())
	}
}
