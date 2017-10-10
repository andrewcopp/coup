package coup

// type Exchange struct {
// 	Subject *Player
// 	Deck    *Deck
// }
//
// func NewExchange(sub *Player, deck *Deck) *Exchange {
// 	return &Exchange{
// 		Subject: sub,
// 		Deck:    deck,
// 	}
// }
//
// func (e *Exchange) Pay() {}
//
// func (e *Exchange) Claim() *Claim {
// 	return NewClaim(e.Subject, Ambassador)
// }
//
// func (e *Exchange) Counter() *func(game *Game) *Block {
// 	return nil
// }
//
// func (e *Exchange) Resolve() func(game *Game) {
// 	return func(game *Game) {
// 		e.Subject.Draw(e.Deck)
// 		e.Subject.Draw(e.Deck)
//
// 		e.Deck.Add(e.Subject.Discard(e.Subject.Chooser.ChooseDiscard()))
// 		e.Deck.Add(e.Subject.Discard(e.Subject.Chooser.ChooseDiscard()))
// 	}
// }
