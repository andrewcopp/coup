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

func (r *Random) ChooseMove(moves []*Move) *Move {
	rand.Seed(int64(time.Now().Nanosecond()))
	return moves[rand.Intn(len(moves))]
}

func (r *Random) ChooseBlock(claims []*Claim) *Claim {
	for _, claim := range claims {
		rand.Seed(int64(time.Now().Nanosecond()))
		if rand.Intn(5) == 0 {
			return claim
		}
	}
	return nil
}

func (r *Random) ChooseChallenge(claim *Claim) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(5) == 0
}

func (r *Random) ChooseDiscard(hand *Cards) CardEnum {
	return hand.Peek()
}
