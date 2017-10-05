package agent

type Block struct {
	Subject   *Player
	Declared  CardType
	Challenge *Challenge
}

type BlockType int

const (
	DukeBlock       BlockType = iota
	AmbassadorBlock           = iota
	CaptainBlock              = iota
	ContessaBlock             = iota
)

func (b *Block) Transition(s *State) *State {
	return s
}
