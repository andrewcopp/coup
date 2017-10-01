package coup

type Move struct {
	Announcement string
	Cost         int
	Claim        *Claim
	Counters     []CardType
	Resolve      func()
}

func NewMove(announcement string, cost int, claim *Claim, counters []CardType, resolve func()) *Move {
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

func (m *Move) Block() *Block {

	// if m.Claim == nil && len(m.Counters) != 0 {
	//
	// }
	//
	// m.Claim.Object.Impede(a.Move.Counters)
	return nil
}
