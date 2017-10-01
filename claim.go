package coup

import "fmt"

type Claim struct {
	Subject   *Player
	Declared  Type
	Object    *Player
	Challenge *Challenge
}

func NewClaim(sub *Player, dec Type, obj *Player) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
		Object:   obj,
	}
}

func (c *Claim) Verify(state *State) {
	challenger := c.Challenge.Subject
	fmt.Printf("%s challenges.\n", challenger.Name)
	revealed := c.Subject.Reveal(state, &c.Declared)
	c.Challenge.Revealed = revealed.Type

	var discarded *Card
	if c.Declared != c.Challenge.Revealed {
		fmt.Printf("Challenge unsuccessful.\n")
		discarded = challenger.Reveal(state, nil)
		state.Deck.Add(revealed)
		c.Subject.Hand = append(c.Subject.Hand, state.Deck.Draw())
	} else {
		fmt.Printf("Challenge successful.\n")
		discarded = revealed
	}

	fmt.Printf("%s discards a %s.\n", c.Subject.Name, discarded.Name())
}
