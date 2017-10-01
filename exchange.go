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
		NewClaim(sub, Ambassador, nil),
		[]Type{},
		exchange.Resolve,
	)
}

func (e *Exchange) Announce() {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)
}

func (e *Exchange) Pay() {}

func (e *Exchange) Resolve(state *State) {
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())
}
