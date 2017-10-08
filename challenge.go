package coup

func NewChallenges(state *State) []*Action {
	return []*Action{}
}

type Challenge struct {
	Subject *Player
}

func NewChallenge(sub *Player) *Challenge {
	return &Challenge{
		Subject: sub,
	}
}
