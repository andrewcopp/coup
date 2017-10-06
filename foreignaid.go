package coup

type ForeignAid struct {
	Subject int
}

func NewForeignAid(sub int) *ForeignAid {
	return &ForeignAid{
		Subject: sub,
	}
}

func (f *ForeignAid) Consider(state *State) []*State {
	return []*State{state}
}

func (f *ForeignAid) Modify(board *Board) {

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
