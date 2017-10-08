package coup

type Coup struct {
	Subject *Player
	Object  *Player
}

func NewCoup(sub *Player, obj *Player) *Coup {
	return &Coup{
		Subject: sub,
		Object:  obj,
	}
}

func (c *Coup) Pay() {
	c.Subject.Coins -= 7
}

func (c *Coup) Claim() *Claim {
	return nil
}

func (c *Coup) Counter() *func(game *Game) *Block {
	return nil
}

func (c *Coup) Resolve() {

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
