package coup

type Decider interface {
	Decide(state *State) *Action
}
