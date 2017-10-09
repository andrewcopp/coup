package coup

type Challenge struct {
	Subject *Player
}

func NewChallenge(sub *Player) *Challenge {
	return &Challenge{
		Subject: sub,
	}
}
