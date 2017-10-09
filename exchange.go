package coup

type Exchange struct {
	Subject *Player
	Deck    *Deck
}

func NewExchange(sub *Player, deck *Deck) *Exchange {
	return &Exchange{
		Subject: sub,
		Deck:    deck,
	}
}

func (e *Exchange) Pay() {}

func (e *Exchange) Claim() *Claim {
	return NewClaim(e.Subject, Ambassador)
}

func (e *Exchange) Counter() *func(game *Game) *Block {
	return nil
}

func (e *Exchange) Resolve() {

}

// func NewExchange(sub *Player, state *State) *Move {
// 	return NewMove(
// 		Exchange,
// 		sub,
// 		fmt.Sprintf("%s exchanges.", sub.Name),
// 		0,
// 		NewClaim(sub, Ambassador, nil),
// 		[]CardType{},
// 		ExchangeFunc(sub, state),
// 	)
// }
//
// func ExchangeFunc(sub *Player, state *State) func() {
// 	return func() {
// 		sub.Draw(state.Deck)
// 		sub.Draw(state.Deck)
//
// 		state.Deck.Add(sub.Reveal(state, nil))
// 		state.Deck.Add(sub.Reveal(state, nil))
// 	}
// }
