package coup

type Board struct {
	Deck    *Deck
	Discard *Hand
}

func NewBoard() *Board {
	return &Board{
		Deck:    NewDeck(),
		Discard: NewHand([]*Card{}),
	}
}
