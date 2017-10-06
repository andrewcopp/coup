package coup

func NewForeignAid(sub int) *Action {
	return &Action{
		StateFunc: func(state *State) []*State {
			return []*State{state}
		}, BoardFunc: func(board *Board) {

		},
	}
}

// func NewForeignAid(sub *Player) *Move {
// 	return NewMove(
// 		ForeignAid,
// 		sub,
// 		fmt.Sprintf("%s takes foreign aid.", sub.Name),
// 		0,
// 		nil,
// 		[]CardType{Duke},
// 		ForeignAidFunc(sub),
// 	)
// }
//
// func ForeignAidFunc(sub *Player) func() {
// 	return func() {
// 		sub.Coins += 2
// 		Account(sub)
// 	}
// }
