package coup

type Action struct {
	Move    *Move
	Counter *Counter
}

func NewAction(mv *Move) *Action {
	return &Action{
		Move:    mv,
		Counter: nil,
	}
}

func (a *Action) Apply(state *State) {
	a.Move.Announce()
	a.Move.Pay()
	if a.Move.Claim(state) {
		a.Move.Resolve(state)
	}
}
