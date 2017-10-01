package coup

type Decider interface {
	Decide(state *State, valid []*Move) *Move
	Dispute(claim *Claim) bool
}
