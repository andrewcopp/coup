package coup

// type Hand struct {
// 	Cards []*Card
// }
//
// func NewHand(cards []*Card) *Hand {
// 	return &Hand{
// 		Cards: cards,
// 	}
// }
//
// func (h *Hand) Size() int {
// 	return len(h.Cards)
// }
//
// func (h *Hand) Add(card *Card) {
// 	h.Cards = append(h.Cards, card)
// }
//
// func (h *Hand) Remove(cardType CardType) *Card {
// 	for i, card := range h.Cards {
// 		if card.CardType == cardType {
// 			h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
// 			return card
// 		}
// 	}
// 	return nil
// }
