package coup

type Tax struct {
	Subject int
}

func NewTax(sub int) *Tax {
	return &Tax{
		Subject: sub,
	}
}

func (t *Tax) Consider(state *State) []*State {
	return []*State{state}
}

func (t *Tax) Modify(board *Board) {

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
