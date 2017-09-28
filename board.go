package coup

type Board struct {
	States []*State
}

func NewBoard() *Board {
	return &Board{
		States: []*State{},
	}
}

func (brd *Board) Setup(players []*Player) {

	cards := make([]*Card, 15)
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			cards[3*i+j] = NewCard(Type(i))
		}
	}

	deck := NewDeck(cards)
	deck.Shuffle()

	for _, player := range players {
		player.Hidden = append(player.Hidden, deck.Draw())
		player.Hidden = append(player.Hidden, deck.Draw())
	}

	brd.States = []*State{NewState(deck, players)}
}

func (brd *Board) Play() *Player {

	for brd.Continue() {
		state := brd.State().Copy()
		player := state.Players[0]
		action := player.Move(state)
		(*action).Modify(state)
		brd.Shift(state)
		brd.States = append(brd.States, state)
	}

	return brd.State().Players[0]
}

func (brd *Board) State() *State {
	return brd.States[len(brd.States)-1]
}

func (brd *Board) Continue() bool {
	return len(brd.State().Alive()) != 1
}

func (brd *Board) Shift(state *State) {
	state.Players = append(state.Players[1:], state.Players[0])

	for len(state.Players[0].Hidden) == 0 {
		state.Players = append(state.Players[1:], state.Players[0])
	}
}
