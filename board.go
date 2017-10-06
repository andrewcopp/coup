package coup

type Board struct {
	Deck    *Hand
	Discard *Hand
}

func NewBoard() *Board {
	return &Board{
		Deck:    NewHand(3, 3, 3, 3, 3),
		Discard: NewHand(0, 0, 0, 0, 0),
	}
}

func (b *Board) Copy() *Board {
	return &Board{
		Deck:    b.Deck.Copy(),
		Discard: b.Discard.Copy(),
	}
}

func (b *Board) Deal(players []*Player) {
	for i := 0; i < 2; i++ {
		for _, player := range append(players[1:], players[0]) {
			player.Draw(b.Deck)
		}
	}
}
