package coup

type Move struct {
	Announcement string
	Cost         int
	Claim        *Claim
	Counters     []Type
	Resolve      func(state *State)
}

func NewMove(announcement string, cost int, claim *Claim, counters []Type, resolve func(state *State)) *Move {
	return &Move{
		Announcement: announcement,
		Cost:         cost,
		Claim:        claim,
		Counters:     counters,
		Resolve:      resolve,
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
