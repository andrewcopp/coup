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

func (r *Random) Decide(state *State) *Move {
	moves := []*Move{}
	self := state.Players[0]

	for _, other := range state.Alive()[1:] {
		if state.Players[0].Coins >= 7 {
			moves = append(moves, NewCoup(self, other))
		}

		if state.Players[0].Coins < 10 {
			moves = append(moves, NewIncome(self))
			moves = append(moves, NewForeignAid(self))
			moves = append(moves, NewTax(self))
			moves = append(moves, NewExchange(self))
			moves = append(moves, NewSteal(self, other))
		}

		if state.Players[0].Coins >= 3 {
			moves = append(moves, NewAssassinate(self, other))
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(moves))
	move := moves[i]
	return move
}

func (r *Random) Dispute(claim *Claim) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(2) != 0 {
		return false
	}

	return true
}

func (r *Random) BlockForeignAid(state *State, sub *Player) *Block {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(2) != 0 {
		return nil
	}

	return &Block{Declared: Duke}
}

func (r *Random) BlockAssassinate(state *State, sub *Player, obj *Player, chg *Player) *Block {
	return nil
}

func (r *Random) BlockSteal(state *State, sub *Player, obj *Player, chg *Player) *Block {
	return nil
}
