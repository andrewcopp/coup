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
	if card := c.Subject.Produce(c.Declared); card != nil {
		c.Challenge.Successful = false
		fmt.Printf("Challenge unsuccessful.\n")
		challenger.Reveal(state)
		fmt.Printf("%s reveals a %s.\n", challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
		state.Deck.Add(card)
		c.Subject.Hand = append(c.Subject.Hand, state.Deck.Draw())
	} else {
		fmt.Printf("Challenge successful.\n")
		c.Subject.Reveal(state)
		c.Challenge.Successful = true
		fmt.Printf("%s reveals a %s.\n", c.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
	}
}
