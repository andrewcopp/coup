package coup

type Move struct {
	Announce func()
	Pay      func()
	Claim    *Claim
	Resolve  func(state *State)
}

func NewMove(announce func(), pay func(), claim *Claim, resolve func(state *State)) *Move {
	return &Move{
		Announce: announce,
		Pay:      pay,
		Claim:    claim,
		Resolve:  resolve,
	}
}

func (m *Move) Successful() bool {
	if m.Claim == nil {
		return true
	}

	return !m.Claim.Challenge.Successful
}

func (m *Move) Scrutinize(state *State) {
	for _, other := range state.Alive()[1:] {
		other.Dispute(m.Claim)
		if m.Claim.Challenge.Subject != nil {
			m.Claim.Verify(state)
			return
		}
	}
}
