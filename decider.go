package coup

type Decider interface {
	Decide(state *State) *Action
	Dispute(claim *Claim) bool

	ChallengeTax(state *State, sub *Player) *Challenge
	ChallengeAssassinate(state *State, sub *Player, obj *Player) *Challenge
	ChallengeExchange(state *State, sub *Player) *Challenge
	ChallengeSteal(state *State, sub *Player, obj *Player) *Challenge

	BlockForeignAid(state *State, sub *Player) *Block
	BlockAssassinate(state *State, sub *Player, obj *Player, chg *Player) *Block
	BlockSteal(state *State, sub *Player, obj *Player, chg *Player) *Block

	ChallengeForeignAidBlock(state *State, sub *Player, blk *Player) *Challenge
	ChallengeAssassinateBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge
	ChallengeStealBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge
}
