package agent

type Action interface {
	Transition(s *State) *State
}
