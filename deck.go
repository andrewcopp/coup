package coup

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []*Card
}

func NewDeck(cards []*Card) *Deck {
	return &Deck{
		Cards: cards,
	}
}

func (d *Deck) Copy() *Deck {
	cards := make([]*Card, len(d.Cards))
	for i, card := range d.Cards {
		cards[i] = card.Copy()
	}
	return NewDeck(cards)
}

type Shuffler interface {
	Shuffle()
}

type Drawer interface {
	Draw() Card
}

func (d *Deck) Shuffle() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) Draw() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) Add(card *Card) {
	d.Cards = append(d.Cards, card)
	d.Shuffle()
}
