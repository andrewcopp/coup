package agent

type Challenge struct {
	Subject    *Player
	Successful bool
	Discarded  CardType
}

type CardType int

const (
	Duke       CardType = iota
	Assassin            = iota
	Ambassador          = iota
	Captain             = iota
	Contessa            = iota
)

func (c *Challenge) Transition(s *State) *State {
	return s
}
