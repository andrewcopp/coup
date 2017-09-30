package coup

import "fmt"

type Exchange struct {
	Subject   *Player
	Challenge *Challenge
}

func NewExchange(sub *Player) *Exchange {
	return &Exchange{
		Subject: sub,
	}
}

func (e *Exchange) Modify(state *State) {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)

	claim := NewClaim(e.Subject, Ambassador, nil)
	for _, other := range state.Alive()[1:] {
		if claim.Challenger == nil {
			other.Dispute(claim)
		}
	}

	if claim.Challenger != nil {
		fmt.Printf("%s challenges.\n", claim.Challenger.Name)
		if card := e.Subject.Produce(claim.Declared); card != nil {
			fmt.Printf("Challenge unsuccessful.\n")
			claim.Challenger.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", claim.Challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
			state.Deck.Add(card)
			e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
		} else {
			fmt.Printf("Challenge successful.\n")
			e.Subject.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", e.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
			return
		}
	}

	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())
}
