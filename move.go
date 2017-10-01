package coup

type Move struct {
	Announce func()
	Pay      func()
	Claim    *Claim
	Counters []Type
	Resolve  func(state *State)
}

func NewMove(announce func(), pay func(), claim *Claim, counters []Type, resolve func(state *State)) *Move {
	return &Move{
		Announce: announce,
		Pay:      pay,
		Claim:    claim,
		Counters: counters,
		Resolve:  resolve,
	}
}

func (m *Move) Successful() bool {
	if m.Claim == nil {
		return true
	}

	if m.Claim.Challenge == nil {
		return true
	}

	return m.Claim.Declared == m.Claim.Challenge.Revealed
}

func (m *Move) Scrutinize(state *State) {
	for _, other := range state.Alive()[1:] {
		other.Dispute(m.Claim)
		if m.Claim.Challenge != nil {
			m.Claim.Verify(state)
			return
		}
	}
}
