package coup

type Decider interface {
	Decide(state *State) *Move
	Dispute(claim *Claim) bool

	BlockForeignAid(state *State, sub *Player) *Block
	BlockAssassinate(state *State, sub *Player, obj *Player, chg *Player) *Block
	BlockSteal(state *State, sub *Player, obj *Player, chg *Player) *Block

	ChallengeForeignAidBlock(state *State, sub *Player, blk *Player) *Challenge
	ChallengeAssassinateBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge
	ChallengeStealBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge
}
