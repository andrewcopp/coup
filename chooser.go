package coup

type Chooser interface {
	ChooseMove(moves []*Move) *Move
	ChooseBlock(claims []*Claim) *Claim
	ChooseChallenge(claim *Claim) bool
	ChooseDiscard(hand *Cards) CardEnum
}
