package coup

import (
	"math/rand"
	"time"
)

type Game struct {
	Board   *Board
	Players []*Player
}

func NewGame(players []*Player) *Game {
	return &Game{
		Board:   NewBoard(),
		Players: players,
	}
}

func (g *Game) Setup() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range g.Players {
		j := rand.Intn(i + 1)
		g.Players[i], g.Players[j] = g.Players[j], g.Players[i]
	}

	for _, player := range g.Players {
		player.Coins = 2
	}

	for i := 0; i < 2; i++ {
		for _, player := range append(g.Players[1:], g.Players[0]) {
			player.Draw(g.Board.Deck)
		}
	}
}

func (g *Game) Play() *Player {

	for g.Next() {
		move := g.Players[0].Move(g)

		successful := true
		if claim := move.Claim(); claim != nil {
			if challenge := g.Dispute(g.Players[0], claim); challenge != nil {
				if g.Verify(claim) {
					g.Board.Discard.Add(challenge.Subject.Discard(challenge.Subject.Chooser.ChooseDiscard()))
					g.Board.Deck.Add(challenge.Subject.Discard(claim.Declared))
					claim.Subject.Draw(g.Board.Deck)
				} else {
					g.Board.Discard.Add(g.Players[0].Discard(g.Players[0].Chooser.ChooseDiscard()))
					successful = false
				}
			}
		}

		// object is stil alive (assassin)
		// discard (exchange)

		if successful {
			if counter := move.Counter(); counter != nil {
				if block := (*counter)(g); block != nil {
					if challenge := g.Dispute(block.Subject, block.Claim); challenge != nil {
						if g.Verify(block.Claim) {
							g.Board.Discard.Add(challenge.Subject.Discard(challenge.Subject.Chooser.ChooseDiscard()))
							g.Board.Deck.Add(block.Claim.Subject.Discard(block.Claim.Declared))
							block.Claim.Subject.Draw(g.Board.Deck)
							successful = false
						} else {
							g.Board.Discard.Add(block.Subject.Discard(block.Subject.Chooser.ChooseDiscard()))
						}
					} else {
						successful = false
					}
				}
			}

			if successful {
				move.Resolve()
			}
		}

	}

	return g.Players[0]
}

func (g *Game) Dispute(sub *Player, claim *Claim) *Challenge {
	for _, opponent := range sub.Opponents(g) {
		if challenge := opponent.Challenge(g, claim); challenge != nil {
			return challenge
		}
	}
	return nil
}

func (g *Game) Verify(claim *Claim) bool {
	return claim.Subject.Reveal(claim.Declared) != nil
}

func (g *Game) Next() bool {
	g.Players = append(g.Players[1:], g.Players[0])
	for g.Players[0].Hand.Size() == 0 {
		g.Players = append(g.Players[1:], g.Players[0])
	}

	for _, player := range g.Players[1:] {
		if player.Alive() {
			return true
		}
	}
	return false
}
