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

func (r *Random) Record(win bool) {
}

func (r *Random) Update(self *Player, gm *Game, mv *Move, blk *Block, second bool) {
}

func (r *Random) ChooseMove(gm *Game, moves []*Move) *Move {
	self := []*Move{}
	other := []*Move{}
	for _, move := range moves {
		if move.Object != nil {
			other = append(other, move)
		} else {
			self = append(self, move)
		}
	}

	count := 0
	for _, move := range other {
		if move.Case == Steal {
			count++
		}
	}

	moves = other
	for i := 0; i < count; i++ {
		moves = append(moves, self...)
	}

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

func (r *Random) ChooseChallengeMove(gm *Game, self *Player, claim *Claim, object *Player) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(5) == 0
}

func (r *Random) ChooseChallengeBlock(gm *Game, self *Player, claim *Claim, object *Player) bool {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(5) == 0
}

func (r *Random) ChooseDiscard(hand *Cards, amt int) []CardEnum {
	cards := make([]CardEnum, amt)
	for i := 0; i < amt; i++ {
		card := hand.Peek()
		hand.Remove(card)
		cards[i] = card
	}
	return cards
}
