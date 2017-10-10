package coup

// type Deck struct {
// 	Cards []*Card
// }
//
// func NewDeck() *Deck {
//
// 	cards := make([]*Card, 15)
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 3; j++ {
// 			cards[3*i+j] = NewCard(CardType(i))
// 		}
// 	}
//
// 	return &Deck{
// 		Cards: cards,
// 	}
// }
//
// type Shuffler interface {
// 	Shuffle()
// }
//
// type Drawer interface {
// 	Draw() Card
// }
//
// func (d *Deck) Shuffle() {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// 	for i := range d.Cards {
// 		j := rand.Intn(i + 1)
// 		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
// 	}
// }
//
// func (d *Deck) Draw() *Card {
// 	card := d.Cards[0]
// 	d.Cards = d.Cards[1:]
// 	return card
// }
//
// func (d *Deck) Add(card *Card) {
// 	d.Cards = append(d.Cards, card)
// 	d.Shuffle()
// }
