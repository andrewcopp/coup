package coup

type State struct {
	Deck     *Deck
	Revealed []*Card
	Players  []*Player
}

func NewState(deck *Deck, revealed []*Card, players []*Player) *State {
	return &State{
		Deck:     deck,
		Revealed: revealed,
		Players:  players,
	}
}

func (s *State) Copy() *State {

	revealed := make([]*Card, len(s.Revealed))
	for i, card := range s.Revealed {
		revealed[i] = card.Copy()
	}

	players := make([]*Player, len(s.Players))
	for i, player := range s.Players {
		players[i] = player.Copy()
	}

	return NewState(
		s.Deck.Copy(),
		revealed,
		players,
	)
}

func (s *State) Alive() []*Player {
	players := []*Player{}
	for _, player := range s.Players {
		if len(player.Hand) != 0 {
			players = append(players, player)
		}
	}
	return players
}
