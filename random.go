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

func (r *Random) Decide(state *State, valid []*Move) *Move {
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(valid))
	move := valid[i]
	return move
}

func (r *Random) Dispute(claim *Claim) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	if rand.Intn(5) != 0 {
		return false
	}

	return true
}

func (r *Random) Discard(state *State) *Card {
	self := state.Players[0]
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(self.Hand))
	card := self.Hand[i]
	self.Hand = append(self.Hand[:i], self.Hand[i+1:]...)
	return card
}
