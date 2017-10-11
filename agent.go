package coup

type Agent struct {
}

func (a *Agent) Update(gm *Game, mv *Move, mvChallenger *Player, blocker *Player, blkChallenger *Player) {
}

func (a *Agent) ChooseMove(moves []*Move) *Move {
	return nil
}

func (a *Agent) ChooseBlock(claims []*Claim) *Claim {
	return nil
}

func (a *Agent) ChooseChallenge(claim *Claim) bool {
	return false
}

func (a *Agent) ChooseDiscard(hand *Cards, amt int) []CardEnum {
	return []CardEnum{}
}
