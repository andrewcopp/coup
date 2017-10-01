package coup

type Challenge struct {
	Subject  *Player
	Revealed Type
}

func NewChallenge(sub *Player) *Challenge {
	return &Challenge{
		Subject: sub,
	}
}
