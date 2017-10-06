package coup

func NewTax(sub int) *Action {
	return &Action{
		StateFunc: func(state *State) []*State {
			return []*State{state}
		}, BoardFunc: func(board *Board) {

		},
	}
}

// func NewTax(sub *Player) *Move {
// 	return NewMove(
// 		Tax,
// 		sub,
// 		fmt.Sprintf("%s taxes.", sub.Name),
// 		0,
// 		NewClaim(sub, Duke, nil),
// 		[]CardType{},
// 		TaxFunc(sub),
// 	)
// }
//
// func TaxFunc(sub *Player) func() {
// 	return func() {
// 		sub.Coins += 3
// 		Account(sub)
// 	}
// }
