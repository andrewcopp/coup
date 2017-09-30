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

func (c *Coup) Announce() {
	fmt.Printf("%s coups %s.\n", c.Subject.Name, c.Object.Name)
}

func (c *Coup) Pay() {
	c.Subject.Coins -= 7
	Account(c.Subject)
}

func (c *Coup) Claim(state *State) bool {
	return true
}

func (c *Coup) Resolve(state *State) {
	c.Object.Reveal(state)
	fmt.Printf("%s reveals a %s.\n", c.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
}

func (c *Coup) Modify(state *State) {
}
