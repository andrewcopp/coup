package coup

import (
	"math/rand"
	"time"
)

type Cards struct {
	Dukes       int
	Assassins   int
	Ambassadors int
	Captains    int
	Contessas   int
}

func NewCards(dukes int, assassins int, ambassadors int, captains int, contessas int) *Cards {
	return &Cards{
		Dukes:       dukes,
		Assassins:   assassins,
		Ambassadors: ambassadors,
		Captains:    captains,
		Contessas:   contessas,
	}
}

func (c *Cards) Size() int {
	return c.Dukes + c.Assassins + c.Ambassadors + c.Captains + c.Contessas
}

func (c *Cards) Add(card CardEnum) {
	switch card {
	case Duke:
		c.Dukes++
	case Assassin:
		c.Assassins++
	case Ambassador:
		c.Ambassadors++
	case Captain:
		c.Captains++
	case Contessa:
		c.Contessas++
	}
}

func (c *Cards) Remove(card CardEnum) {
	switch card {
	case Duke:
		c.Dukes--
	case Assassin:
		c.Assassins--
	case Ambassador:
		c.Ambassadors--
	case Captain:
		c.Captains--
	case Contessa:
		c.Contessas--
	}
}

func (c *Cards) Peek() CardEnum {
	cards := []CardEnum{}
	for i := 0; i < c.Dukes; i++ {
		cards = append(cards, Duke)
	}

	for i := 0; i < c.Assassins; i++ {
		cards = append(cards, Assassin)
	}

	for i := 0; i < c.Dukes; i++ {
		cards = append(cards, Ambassador)
	}

	for i := 0; i < c.Dukes; i++ {
		cards = append(cards, Captain)
	}

	for i := 0; i < c.Dukes; i++ {
		cards = append(cards, Contessa)
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	return cards[rand.Intn(len(cards))]
}
