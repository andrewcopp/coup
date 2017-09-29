package coup

import "fmt"

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

func (c *Coup) Modify(state *State) {
	c.Subject.Coins -= 7
	c.Object.Reveal(state)
	fmt.Printf("%s coups %s.\n", c.Subject.Name, c.Object.Name)
	Account(c.Subject)
	fmt.Printf("%s reveals a %s.\n", c.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
}

func (c *Coup) Dispute(state *State) {}

func (c *Coup) Impede(state *State) {}

func (c *Coup) Describe() {
	// fmt.Printf("%s coups %s.\n", c.Subject.Name, c.Object.Name)
	// Account(c.Subject)
	// fmt.Printf("%s reveals a %s.\n", c.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
}
