package coup

import "fmt"

type Coup struct {
	Subject *Player
	Object  *Player
}

func NewCoup(sub *Player, obj *Player) *Move {
	coup := Coup{
		Subject: sub,
		Object:  obj,
	}

	return NewMove(
		fmt.Sprintf("%s coups %s.", sub.Name, obj.Name),
		7,
		nil,
		[]Type{},
		coup.Resolve,
	)
}

func (c *Coup) Resolve(state *State) {
	revealed := c.Object.Reveal(state, nil)
	state.Revealed = append(state.Revealed, revealed)
	fmt.Printf("%s reveals a %s.\n", c.Object.Name, revealed.Name())
}
