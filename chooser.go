package coup

// The interface required to make decisions in the game.
// Implemented by Random, Human, and Agent.

type Chooser interface {
	Record(win bool)
	Update(self *Player, gm *Game, mv *Move, blk *Block, second bool)
	ChooseMove(gm *Game, moves []*Move) *Move
	ChooseBlock(claims []*Claim) *Claim
	ChooseChallengeMove(gm *Game, self *Player, claim *Claim, object *Player) bool
	ChooseChallengeBlock(gm *Game, self *Player, claim *Claim, object *Player) bool
	ChooseDiscard(hand *Cards, amt int) []CardEnum
}
