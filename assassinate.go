package coup

func NewAssassinate(sub int, obj int) *Action {
	return &Action{
		StateFunc: func(state *State) []*State {
			return []*State{state}
		}, BoardFunc: func(board *Board) {

		},
	}
}

// func NewAssassinate(sub *Player, obj *Player, state *State) *Move {
// 	return NewMove(
// 		Assassinate,
// 		sub,
// 		fmt.Sprintf("%s assassinates %s.", sub.Name, obj.Name),
// 		3,
// 		NewClaim(sub, Assassin, obj),
// 		[]CardType{Contessa},
// 		AssassinateFunc(sub, obj, state),
// 	)
// }
//
// func AssassinateFunc(sub *Player, obj *Player, state *State) func() {
// 	return func() {
// 		if len(obj.Hand) != 0 {
// 			revealed := obj.Reveal(state, nil)
// 			state.Discard.Add(revealed)
// 			fmt.Printf("%s discards a %s.\n", obj.Name, revealed.Name())
// 		}
// 	}
// }
