package coup

type Hand struct {
	Cards []*Card
}

func NewHand(cards []*Card) *Hand {
	return &Hand{
		Cards: cards,
	}
}

func (h *Hand) Copy() *Hand {
	cards := make([]*Card, len(h.Cards))
	for i, card := range h.Cards {
		cards[i] = card
	}
	return NewHand(cards)
}

func (h *Hand) Size() int {
	return len(h.Cards)
}

func (h *Hand) Add(card *Card) {
	h.Cards = append(h.Cards, card)
}
