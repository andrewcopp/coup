package coup

type Claim struct {
	Subject  *Player
	Declared CardEnum
}

func NewClaim(sub *Player, dec CardEnum) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
	}
}

func (c *Claim) Verify() bool {
	switch c.Declared {
	case Duke:
		return c.Subject.Hand.Dukes > 0
	case Assassin:
		return c.Subject.Hand.Assassins > 0
	case Ambassador:
		return c.Subject.Hand.Ambassadors > 0
	case Captain:
		return c.Subject.Hand.Captains > 0
	case Contessa:
		return c.Subject.Hand.Contessas > 0
	}

	panic("Non-Existant Card")
}
