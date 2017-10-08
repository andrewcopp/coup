package coup

type CardType int

const (
	Duke       CardType = iota
	Ambassador          = iota
	Assassin            = iota
	Captain             = iota
	Contessa            = iota
)

// type Card struct {
// 	CardType CardType
// }

// func NewCard(c CardType) *Card {
// 	return &Card{
// 		CardType: c,
// 	}
// }
//
// func (c *Card) Copy() *Card {
// 	return NewCard(c.CardType)
// }
//
// func (c *Card) Name() string {
// 	switch c.CardType {
// 	case Duke:
// 		return "Duke"
// 	case Ambassador:
// 		return "Ambassador"
// 	case Assassin:
// 		return "Assassin"
// 	case Captain:
// 		return "Captain"
// 	case Contessa:
// 		return "Contessa"
// 	}
//
// 	panic("Invalid Card.")
// }
