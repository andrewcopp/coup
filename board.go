package coup

import (
	"math/rand"
	"time"
)

type Board struct {
	Deck    *Hand
	Discard *Hand
	Players []*Player
}

func NewBoard() *Board {
	return &Board{
		Deck:    NewHand(3, 3, 3, 3, 3),
		Discard: NewHand(0, 0, 0, 0, 0),
		Players: []*Player{},
	}
}

func (b *Board) Copy() *Board {
	players := make([]*Player, len(b.Players))
	for i, player := range b.Players {
		players[i] = player.Copy()
	}

	return &Board{
		Deck:    b.Deck.Copy(),
		Discard: b.Discard.Copy(),
		Players: players,
	}
}

func (b *Board) Add(player *Player) {
	player.Coins = 2
	b.Players = append(b.Players, player)
}

func (b *Board) Setup() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range b.Players {
		j := rand.Intn(i + 1)
		b.Players[i], b.Players[j] = b.Players[j], b.Players[i]
	}
}

func (b *Board) Deal() {
	for i := 0; i < 2; i++ {
		for _, player := range append(b.Players[1:], b.Players[0]) {
			player.Draw(b.Deck)
		}
	}
}

func (b *Board) Play() *Player {

	for b.Next() {
		b.Move()
		b.Challenge()
		b.Block()
		b.Challenge()
	}

	return b.Players[0]
}

func (b *Board) Next() bool {
	b.Players = append(b.Players[1:], b.Players[0])
	for b.Players[0].Hand.Size() == 0 {
		b.Players = append(b.Players[1:], b.Players[0])
	}

	for _, player := range b.Players[1:] {
		if player.Alive() {
			return true
		}
	}
	return false
}

func (b *Board) State(player int) *State {
	players := append(b.Players[player:], b.Players[:player]...)
	self := NewSelf(players[0])
	others := make([]*Other, len(players[1:]))
	for i, player := range players[1:] {
		others[i] = NewOther(player)
	}
	return NewState(self, others, b.Discard)
}

func (b *Board) Move() {
	state := b.State(0)
	actions := b.Players[0].Moves(state)
	print(actions)
}

func (b *Board) Challenge() {

}

func (b *Board) Block() {

}
