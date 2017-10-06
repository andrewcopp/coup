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

func (g *Game) Add(player *Player) {
	player.Coins = 2
	g.Players = append(g.Players, player)
}

func (g *Game) Setup() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range g.Players {
		j := rand.Intn(i + 1)
		g.Players[i], g.Players[j] = g.Players[j], g.Players[i]
	}
}

func (g *Game) Play() *Player {

	for g.Next() {
		g.Move()
		g.Challenge()
		g.Block()
		g.Challenge()
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

func (g *Game) State(player int) *State {
	players := append(g.Players[player:], g.Players[:player]...)
	self := NewSelf(players[0])
	others := make([]*Other, len(players[1:]))
	for i, player := range players[1:] {
		others[i] = NewOther(player)
	}
	return NewState(self, others, g.Board.Discard)
}

func (g *Game) Move() {
	state := g.State(0)
	actions := g.Players[0].Moves(state)
	print(actions)
}

func (g *Game) Challenge() {

}

func (g *Game) Block() {

}
