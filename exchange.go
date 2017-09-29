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
	e.Subject.Hidden = append(e.Subject.Hidden, state.Deck.Draw())
	e.Subject.Hidden = append(e.Subject.Hidden, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())
}

func (e *Exchange) Dispute() {

}

func (e *Exchange) Impede() {}

func (e *Exchange) Describe() {
	fmt.Printf("%s exchanges.\n", e.Subject.Name)
}
