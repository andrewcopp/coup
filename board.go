package coup

import "fmt"

type Board struct {
	State *State
}

func NewBoard() *Board {
	return &Board{
		State: nil,
	}
}

func (brd *Board) Setup(players []*Player) {

	cards := make([]*Card, 15)
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			cards[3*i+j] = NewCard(CardType(i))
		}
	}

	deck := NewDeck(cards)
	deck.Shuffle()

	for _, player := range players {
		player.Draw(deck)
		player.Draw(deck)
	}

	brd.State = NewState(nil, deck, NewPile(), players)
}

func (brd *Board) Play() *Player {

	for count := 1; brd.Continue(); count++ {
		fmt.Println()
		fmt.Printf("Turn %d\n", count)
		fmt.Println("------")
		state := brd.State.Copy()
		player := state.Players[0]
		action := player.Move(state)
		action.Apply(state)
		brd.Shift(state)
		brd.State = state
	}

	return brd.State.Players[0]
}

func (brd *Board) Continue() bool {
	return len(brd.State.Alive()) != 1
}

func (brd *Board) Shift(state *State) {
	state.Players = append(state.Players[1:], state.Players[0])

	for len(state.Players[0].Hand) == 0 {
		state.Players = append(state.Players[1:], state.Players[0])
	}
}
