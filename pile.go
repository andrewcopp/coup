package coup

type Pile struct {
	Dukes       []*Card
	Assassins   []*Card
	Ambassadors []*Card
	Captains    []*Card
	Contessas   []*Card
}

func NewPile() *Pile {
	return &Pile{
		Dukes:       []*Card{},
		Assassins:   []*Card{},
		Ambassadors: []*Card{},
		Captains:    []*Card{},
		Contessas:   []*Card{},
	}
}

func (p *Pile) Copy() *Pile {
	pile := NewPile()

	dukes := make([]*Card, len(p.Dukes))
	for i, card := range p.Dukes {
		dukes[i] = card
	}
	pile.Dukes = dukes

	assassins := make([]*Card, len(p.Assassins))
	for i, card := range p.Assassins {
		assassins[i] = card
	}
	pile.Assassins = assassins

	ambassadors := make([]*Card, len(p.Ambassadors))
	for i, card := range p.Ambassadors {
		ambassadors[i] = card
	}
	pile.Ambassadors = ambassadors

	captains := make([]*Card, len(p.Captains))
	for i, card := range p.Captains {
		captains[i] = card
	}
	pile.Captains = captains

	contessas := make([]*Card, len(p.Contessas))
	for i, card := range p.Contessas {
		contessas[i] = card
	}
	pile.Contessas = contessas

	return pile
}

func (p *Pile) Add(c *Card) {
	switch c.CardType {
	case Duke:
		p.Dukes = append(p.Dukes, c)
	case Assassin:
		p.Assassins = append(p.Assassins, c)
	case Ambassador:
		p.Ambassadors = append(p.Ambassadors, c)
	case Captain:
		p.Captains = append(p.Captains, c)
	case Contessa:
		p.Contessas = append(p.Contessas, c)
	}
}
