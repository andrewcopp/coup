package coup

func NewIncome(sub int) *Action {
	return &Action{
		StateFunc: func(state *State) []*State {
			return []*State{state}
		}, BoardFunc: func(board *Board) {
			// board.Players[sub].Coins++
		},
	}
}

// func NewIncome(sub *Player) *Move {
// 	return NewMove(
// 		Income,
// 		sub,
// 		fmt.Sprintf("%s takes income.", sub.Name),
// 		0,
// 		nil,
// 		[]CardType{},
// 		IncomeFunc(sub),
// 	)
// }
//
// func IncomeFunc(sub *Player) func() {
// 	return func() {
// 		sub.Coins++
// 		Account(sub)
// 	}
// }
