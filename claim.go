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

	var loser *Player
	var discarded *Card
	if c.Declared != c.Challenge.Revealed {
		fmt.Printf("Challenge successful.\n")
		loser = c.Subject
		discarded = revealed
	} else {
		fmt.Printf("Challenge unsuccessful.\n")
		loser = challenger
		discarded = challenger.Reveal(state, nil)
		state.Deck.Add(revealed)
		c.Subject.Hand = append(c.Subject.Hand, state.Deck.Draw())
	}

	fmt.Printf("%s discards a %s.\n", loser.Name, discarded.Name())
}
