package coup

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
