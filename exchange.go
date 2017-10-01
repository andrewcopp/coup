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
		[]CardType{},
		exchange.Resolve,
	)
}

func (e *Exchange) Resolve(state *State) {
	e.Subject.Draw(state.Deck)
	e.Subject.Draw(state.Deck)

	state.Deck.Add(e.Subject.Reveal(state, nil))
	state.Deck.Add(e.Subject.Reveal(state, nil))
}
