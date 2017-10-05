package agent

type Move struct {
	Type      MoveType
	Subject   *Player
	Object    *Player
	Challenge *Challenge
}

type MoveType struct {
	Income      bool
	ForeignAid  bool
	Coup        bool
	Tax         bool
	Assassinate bool
	Exchange    bool
	Steal       bool
}

func (m *Move) Transition(s *State) *State {
	return s
}
