package coup

type Chooser interface {
	Update(gm *Game, mv *Move, mvChallenger *Player, blocker *Player, blkChallenger *Player)
	ChooseMove(moves []*Move) *Move
	ChooseBlock(claims []*Claim) *Claim
	ChooseChallenge(claim *Claim) bool
	ChooseDiscard(hand *Cards, amt int) []CardEnum
}
