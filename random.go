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

func (r *Random) ChooseMove(moves []Move) Move {
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

func (r *Random) ChooseDiscard() CardType {
	return Duke
}

// func (r *Random) Dispute(claim *Claim) bool {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// 	if rand.Intn(5) != 0 {
// 		return false
// 	}
//
// 	return true
// }
//
// func (r *Random) Impede(counter CardType) bool {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// 	if rand.Intn(5) != 0 {
// 		return false
// 	}
//
// 	return true
// }
//
// func (r *Random) Discard(state *State, player *Player) *Card {
// 	rand.Seed(int64(time.Now().Nanosecond()))
// 	i := rand.Intn(len(player.Hand))
// 	card := player.Hand[i]
// 	player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
// 	return card
// }
