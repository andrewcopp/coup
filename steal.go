package coup

type Steal struct {
	Subject int
	Object  int
}

func NewSteal(sub int, obj int) *Steal {
	return &Steal{
		Subject: sub,
		Object:  obj,
	}
}

func (s *Steal) Consider(state *State) []*State {
	return []*State{state}
}

func (s *Steal) Modify(board *Board) {

}

// func NewSteal(sub *Player, obj *Player) *Move {
// 	return NewMove(
// 		Steal,
// 		sub,
// 		fmt.Sprintf("%s steals from %s.", sub.Name, obj.Name),
// 		0,
// 		NewClaim(sub, Captain, obj),
// 		[]CardType{Ambassador, Captain},
// 		StealFunc(sub, obj),
// 	)
// }
//
// func StealFunc(sub *Player, obj *Player) func() {
// 	return func() {
// 		amt := 2
// 		if obj.Coins < 2 {
// 			amt = obj.Coins
// 		}
//
// 		obj.Coins -= amt
// 		sub.Coins += amt
//
// 		Account(sub)
// 		Account(obj)
// 	}
// }
