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

	for _, other := range state.Alive()[1:] {
		if state.Players[0].Coins >= 7 {
			var coup Action
			coup = NewCoup(state.Players[0], other)
			actions = append(actions, &coup)
		}

		if state.Players[0].Coins < 10 {
			var income Action
			income = NewIncome(state.Players[0])
			actions = append(actions, &income)

			var foreignAid Action
			foreignAid = NewForeignAid(state.Players[0])
			actions = append(actions, &foreignAid)

			var tax Action
			tax = NewTax(state.Players[0])
			actions = append(actions, &tax)

			var exchange Action
			exchange = NewExchange(state.Players[0])
			actions = append(actions, &exchange)

			var steal Action
			steal = NewSteal(state.Players[0], other)
			actions = append(actions, &steal)
		}

		if state.Players[0].Coins >= 3 {
			var assassinate Action
			assassinate = NewAssassinate(state.Players[0], other)
			actions = append(actions, &assassinate)
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
