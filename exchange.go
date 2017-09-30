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

func (e *Exchange) Announce() {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)
}

func (e *Exchange) Modify(state *State) {

	claim := NewClaim(e.Subject, Ambassador, nil)
	e.Subject.Make(claim, state)
	if claim.Revealed == nil {
		e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
		e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

		state.Deck.Add(e.Subject.Discard())
		state.Deck.Add(e.Subject.Discard())
		return
	}

	if *claim.Revealed == claim.Declared {
		e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
		e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

		state.Deck.Add(e.Subject.Discard())
		state.Deck.Add(e.Subject.Discard())
		return
	}
}
