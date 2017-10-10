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
	i := rand.Intn(len(moves))
	move := moves[i]
	return move
}

func (r *Random) ChooseBlock(claim *Claim) bool {
	return false
}

func (r *Random) ChooseChallenge(claim *Claim) bool {
	return false
}

func (r *Random) ChooseDiscard() CardEnum {
	return Duke
}
