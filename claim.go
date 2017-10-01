package coup

import "fmt"

type Claim struct {
	Subject   *Player
	Declared  CardType
	Object    *Player
	Challenge *Challenge
}

func NewClaim(sub *Player, dec CardType, obj *Player) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
		Object:   obj,
	}
}

func (c *Claim) Name() string {
	switch c.Declared {
	case Duke:
		return "Duke"
	case Assassin:
		return "Assassin"
	case Ambassador:
		return "Ambassador"
	case Captain:
		return "Captain"
	case Contessa:
		return "Contessa"
	}

	panic("Weird CardType")
}

func (c *Claim) Scrutinize(state *State) {
	for _, other := range c.Subject.Opponents(state) {
		other.Dispute(c)
		if c.Challenge != nil {
			c.Verify(state)
			return
		}
	}
}

func (c *Claim) Verify(state *State) {
	challenger := c.Challenge.Subject
	fmt.Printf("%s challenges.\n", challenger.Name)
	revealed := c.Subject.Reveal(state, &c.Declared)
	c.Challenge.Revealed = revealed.CardType

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
		c.Subject.Draw(state.Deck)
	}

	fmt.Printf("%s discards a %s.\n", loser.Name, discarded.Name())
}
