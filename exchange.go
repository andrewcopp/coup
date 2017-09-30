package coup

import "fmt"

type Exchange struct {
	Subject   *Player
	Challenge *Challenge
}

func NewExchange(sub *Player) *Move {
	exchange := Exchange{
		Subject: sub,
	}

	return NewMove(
		exchange.Announce,
		exchange.Pay,
		exchange.Claim,
		exchange.Resolve,
	)
}

func (e *Exchange) Announce() {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)
}

func (e *Exchange) Pay() {}

func (e *Exchange) Claim(state *State) bool {
	claim := NewClaim(e.Subject, Ambassador, nil)
	e.Subject.Make(claim, state)
	if claim.Revealed == nil {
		return true
	}

	if *claim.Revealed == claim.Declared {
		return true
	}

	return false
}

func (e *Exchange) Resolve(state *State) {
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())
}

func (e *Exchange) Modify(state *State) {

}
