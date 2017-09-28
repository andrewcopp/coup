package coup

type State struct {
	Deck    *Deck
	Players []*Player
}

func NewState(deck *Deck, players []*Player) *State {
	return &State{
		Deck:    deck,
		Players: players,
	}
}

func (s *State) Copy() *State {

	players := make([]*Player, len(s.Players))
	for i, player := range s.Players {
		players[i] = player.Copy()
	}

	return NewState(
		s.Deck.Copy(),
		players,
	)
}

func (s *State) Alive() []*Player {
	players := []*Player{}
	for _, player := range s.Players {
		if len(player.Hidden) != 0 {
			players = append(players, player)
		}
	}
	return players
}
