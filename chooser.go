package coup

type Chooser interface {
	ChooseMove(moves []*Move) *Move
	ChooseBlock(claim *Claim) bool
	ChooseChallenge(claim *Claim) bool
	// ChooseDiscard() CardType
}
