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
	var action Action
	if state.Players[0].Coins >= 7 {
		others := state.Alive()[1:]
		rand.Seed(int64(time.Now().Nanosecond()))
		other := others[rand.Intn(len(others))]
		action = NewOverthrow(state.Players[0], other)
	} else {
		action = NewCollect(state.Players[0])
	}

	return &action
}
