package coup

type Claim struct {
	Subject  *Player
	Declared CardType
}

func NewClaim(sub *Player, dec CardType) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
	}
}
