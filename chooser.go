package coup

type Chooser interface {
	Update(self *Player, gm *Game, mv *Move, blk *Block, second bool)
	ChooseMove(gm *Game, moves []*Move) *Move
	ChooseBlock(claims []*Claim) *Claim
	ChooseChallenge(claim *Claim) bool
	ChooseDiscard(hand *Cards, amt int) []CardEnum
}
