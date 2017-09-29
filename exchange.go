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
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())
	e.Subject.Hand = append(e.Subject.Hand, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())
}

func (e *Exchange) Dispute(state *State) {
	_ = NewClaim(e.Subject, Ambassador, nil)
}

func (e *Exchange) Impede(state *State) {}

func (e *Exchange) Describe() {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)
}
