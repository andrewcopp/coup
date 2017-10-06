package coup

type Income struct {
	Subject int
}

func NewIncome(sub int) *Income {
	return &Income{
		Subject: sub,
	}
}

func (i *Income) Consider(state *State) []*State {
	return []*State{state}
}

func (i *Income) Modify(board *Board) {

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
