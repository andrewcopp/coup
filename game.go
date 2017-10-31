package coup

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Logs    bool
	Deck    *Cards
	Discard *Cards
	Players []*Player
}

func NewGame(players []*Player) *Game {
	return &Game{
		Logs:    false,
		Deck:    NewCards(3, 3, 3, 3, 3),
		Discard: NewCards(0, 0, 0, 0, 0),
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
		if !player.Placeholder {
			player.Coins = 2
		}
	}

	for i := 0; i < 2; i++ {
		for _, player := range append(g.Players[1:], g.Players[0]) {
			if !player.Placeholder {
				card := g.Deck.Peek()
				g.Deck.Remove(card)
				player.Hand.Add(card)
			}
		}
	}
}

func (g *Game) Play() *Player {

	for _, player := range g.Players {
		player.Observe(g, nil, nil, false)
	}

	for i := 1; g.Next(); i++ {
		if g.Logs {
			fmt.Println()
			fmt.Printf("Turn %d\n", i)
			fmt.Println("------------")
			fmt.Println()
		}
		g.Players[0].Move(g).Modify(g)
	}

	return g.Players[0]
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
