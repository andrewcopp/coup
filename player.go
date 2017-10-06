package coup

import (
	"math/rand"
	"time"
)

type Player struct {
	Name string
	// Brain *Decider
	Coins int
	Hand  *Hand
}

func (p *Player) Copy() *Player {
	return &Player{
		Name:  p.Name,
		Coins: p.Coins,
		Hand:  p.Hand.Copy(),
	}
}

//
// func NewPlayer(name string, brain *Decider, coins int) *Player {
// 	return &Player{
// 		Name: name,
// 		// Brain: brain,
// 		Coins: coins,
// 		Hand:  []*Card{},
// 	}
// }
//
// func (p *Player) Copy() *Player {
// 	player := NewPlayer(p.Name, p.Brain, p.Coins)
//
// 	hand := make([]*Card, len(p.Hand))
// 	for i, card := range p.Hand {
// 		hand[i] = card.Copy()
// 	}
// 	player.Hand = hand
//
// 	return player
// }
//
// func (p *Player) Opponents(state *State) []*Player {
// 	opponents := []*Player{}
// 	for _, player := range state.Alive() {
// 		if player != p {
// 			opponents = append(opponents, player)
// 		}
// 	}
// 	return opponents
// }

func (p *Player) Alive() bool {
	return p.Hand.Size() > 0
}

func (p *Player) Draw(deck *Hand) {
	draws := []func(){}
	for i := 0; i < deck.Dukes; i++ {
		draws = append(draws, func() {
			deck.Dukes--
			p.Hand.Dukes++
		})
	}
	for i := 0; i < deck.Assassins; i++ {
		draws = append(draws, func() {
			deck.Assassins--
			p.Hand.Assassins++
		})
	}
	for i := 0; i < deck.Ambassadors; i++ {
		draws = append(draws, func() {
			deck.Ambassadors--
			p.Hand.Ambassadors++
		})
	}
	for i := 0; i < deck.Captains; i++ {
		draws = append(draws, func() {
			deck.Captains--
			p.Hand.Captains++
		})
	}
	for i := 0; i < deck.Contessas; i++ {
		draws = append(draws, func() {
			deck.Contessas--
			p.Hand.Contessas++
		})
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	draws[rand.Intn(len(draws))]()
}

// func (p *Player) Move(state *State) *Action {
// 	valid := []*Move{}
// 	self := state.Players[0]
//
// 	for _, other := range p.Opponents(state) {
// 		if state.Players[0].Coins >= 7 {
// 			valid = append(valid, NewCoup(self, other, state))
// 		}
//
// 		if state.Players[0].Coins < 10 {
// 			valid = append(valid, NewIncome(self))
// 			valid = append(valid, NewForeignAid(self))
// 			valid = append(valid, NewTax(self))
// 			valid = append(valid, NewExchange(self, state))
// 			valid = append(valid, NewSteal(self, other))
// 		}
//
// 		if state.Players[0].Coins >= 3 {
// 			valid = append(valid, NewAssassinate(self, other, state))
// 		}
// 	}
//
// 	move := (*p.Brain).Decide(state, valid)
// 	return NewAction(move)
// }
//
// func (p *Player) Reveal(state *State, t *CardType) *Card {
// 	if t != nil {
// 		for i, card := range p.Hand {
// 			if card.CardType == *t {
// 				p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
// 				return card
// 			}
// 		}
// 	}
//
// 	return (*p.Brain).Discard(state, p)
// }
//
// func (p *Player) Dispute(claim *Claim) {
// 	if (*p.Brain).Dispute(claim) {
// 		claim.Challenge = NewChallenge(p)
// 	}
// }
//
// func (p *Player) Impede(counters []CardType) *Block {
// 	for _, counter := range counters {
// 		if (*p.Brain).Impede(counter) {
// 			return NewBlock(p, counter)
// 		}
// 	}
//
// 	return nil
// }
