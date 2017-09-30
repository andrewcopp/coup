package coup

import (
	"math/rand"
	"time"
)

type Random struct {
}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) Decide(state *State) *Action {
	actions := []*Action{}
	self := state.Players[0]

	for _, other := range state.Alive()[1:] {
		if state.Players[0].Coins >= 7 {
			actions = append(actions, NewCoup(self, other))
		}

		if state.Players[0].Coins < 10 {
			actions = append(actions, NewIncome(self))
			actions = append(actions, NewForeignAid(self))
			actions = append(actions, NewTax(self))
			actions = append(actions, NewExchange(self))
			actions = append(actions, NewSteal(self, other))
		}

		if state.Players[0].Coins >= 3 {
			actions = append(actions, NewAssassinate(self, other))
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(actions))
	action := actions[i]
	return action
}

func (r *Random) Dispute(claim *Claim) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(2) != 0 {
		return false
	}

	return true
}

func (r *Random) ChallengeTax(state *State, player *Player) *Challenge {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(2) != 0 {
		return nil
	}

	return NewChallenge(player)
}

func (r *Random) BlockForeignAid(state *State, sub *Player) *Block {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(2) != 0 {
		return nil
	}

	return NewBlock()
}

func (r *Random) BlockAssassinate(state *State, sub *Player, obj *Player, chg *Player) *Block {
	return nil
}

func (r *Random) BlockSteal(state *State, sub *Player, obj *Player, chg *Player) *Block {
	return nil
}

func (r *Random) ChallengeForeignAidBlock(state *State, sub *Player, blk *Player) *Challenge {
	return nil
}

func (r *Random) ChallengeAssassinateBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge {
	return nil
}

func (r *Random) ChallengeStealBlock(state *State, sub *Player, obj *Player, chg *Player, blk *Player) *Challenge {
	return nil
}
