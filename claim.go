package coup

type Claim struct {
	Declared CardType
	Object   *Player
}

func NewClaim(dec CardType, obj *Player) *Claim {
	return &Claim{
		Declared: dec,
		Object:   obj,
	}
}
