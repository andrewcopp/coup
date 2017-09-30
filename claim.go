package coup

type Claim struct {
	Subject   *Player
	Declared  Type
	Object    *Player
	Challenge *Challenge
}

func NewClaim(sub *Player, dec Type, obj *Player) *Claim {
	return &Claim{
		Subject:  sub,
		Declared: dec,
		Object:   obj,
	}
}
