package coup

type Coup struct {
	Subject int
	Object  int
}

func NewCoup(sub int, obj int) *Coup {
	return &Coup{
		Subject: sub,
		Object:  obj,
	}
}

func (c *Coup) Consider(state *State) []*State {
	return []*State{state}
}

func (c *Coup) Modify(board *Board) {

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
