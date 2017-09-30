package coup

type Challenge struct {
	Subject    *Player
	Successful bool
}

func NewChallenge(sub *Player) *Challenge {
	return &Challenge{
		Subject: sub,
	}
}
