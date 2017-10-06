package coup

type Assassinate struct {
	Subject int
	Object  int
}

func NewAssassinate(sub int, obj int) *Assassinate {
	return &Assassinate{
		Subject: sub,
		Object:  obj,
	}
}

func (a *Assassinate) Consider(state *State) []*State {
	return []*State{state}
}

func (a *Assassinate) Modify(board *Board) {

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
