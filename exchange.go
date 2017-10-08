package coup

type Exchange struct {
	Subject *Player
}

func NewExchange(sub *Player) *Exchange {
	return &Exchange{
		Subject: sub,
	}
}

func (e *Exchange) Pay() {}

func (e *Exchange) Claim() *Claim {
	return NewClaim(Ambassador, nil)
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
