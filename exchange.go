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
		fmt.Sprintf("%s exchanges.", sub.Name),
		0,
		NewClaim(sub, Ambassador, nil),
		[]Type{},
		exchange.Resolve,
	)
}

func (e *Exchange) Announce() {

}

func (e *Exchange) Pay() {}

func (e *Exchange) Resolve(state *State) {
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

	state.Deck.Add(e.Subject.Reveal(state, nil))
	state.Deck.Add(e.Subject.Reveal(state, nil))
}
