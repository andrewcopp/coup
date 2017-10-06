package coup

type Exchange struct {
	Subject int
}

func NewExchange(sub int) *Exchange {
	return &Exchange{
		Subject: sub,
	}
}

func (e *Exchange) Consider(state *State) []*State {
	return []*State{state}
}

func (e *Exchange) Modify(board *Board) {

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
