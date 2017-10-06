package coup

func NewCoup(sub int, obj int) *Action {
	return &Action{
		StateFunc: func(state *State) []*State {
			return []*State{state}
		}, BoardFunc: func(board *Board) {

		},
	}
}

// func NewCoup(sub *Player, obj *Player, state *State) *Move {
// 	return NewMove(
// 		Coup,
// 		sub,
// 		fmt.Sprintf("%s coups %s.", sub.Name, obj.Name),
// 		7,
// 		nil,
// 		[]CardType{},
// 		CoupFunc(sub, obj, state),
// 	)
// }
//
// func CoupFunc(sub *Player, obj *Player, state *State) func() {
// 	return func() {
// 		revealed := obj.Reveal(state, nil)
// 		// state.Discard.Add(revealed)
// 		fmt.Printf("%s reveals a %s.\n", obj.Name, revealed.Name())
// 	}
// }
