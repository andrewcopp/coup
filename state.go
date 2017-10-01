package coup

type State struct {
	Previous *State
	Deck     *Deck
	Discard  *Pile
	Players  []*Player
}

func NewState(prv *State, deck *Deck, discard *Pile, players []*Player) *State {
	return &State{
		Previous: prv,
		Deck:     deck,
		Discard:  discard,
		Players:  players,
	}
}

func (s *State) Copy() *State {

	players := make([]*Player, len(s.Players))
	for i, player := range s.Players {
		players[i] = player.Copy()
	}

	return NewState(
		s,
		s.Deck.Copy(),
		s.Discard.Copy(),
		players,
	)
}

func (s *State) Alive() []*Player {
	players := []*Player{}
	for _, player := range s.Players {
		if player.Alive() {
			players = append(players, player)
		}
	}
	return players
}
