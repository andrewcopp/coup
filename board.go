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

// func (b *Board) Setup() {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// 	for i := range b.Players {
// 		j := rand.Intn(i + 1)
// 		b.Players[i], b.Players[j] = b.Players[j], b.Players[i]
// 	}
// }

// func (b *Board) Deal() {
// 	for i := 0; i < 2; i++ {
// 		for _, player := range append(b.Players[1:], b.Players[0]) {
// 			player.Draw(b.Deck)
// 		}
// 	}
// }
