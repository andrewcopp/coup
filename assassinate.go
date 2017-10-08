package coup

type Assassinate struct {
	Subject *Player
	Object  *Player
}

func NewAssassinate(sub *Player, obj *Player) *Assassinate {
	return &Assassinate{
		Subject: sub,
		Object:  obj,
	}
}

func (a *Assassinate) Pay() {
	a.Subject.Coins -= 3
}

func (a *Assassinate) Claim() *Claim {
	return NewClaim(Assassin, a.Object)
}

func (a *Assassinate) Counter() *func(game *Game) *Block {
	blockFunc := func(game *Game) *Block {
		return nil
	}
	return &blockFunc
}

func (a *Assassinate) Resolve() {

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
