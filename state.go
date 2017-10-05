package coup

type State struct {
	Self            *Self
	Others          []*Other
	Discard         *Hand
	MoveSubjectOne  bool
	MoveSubjectTwo  bool
	MoveTax         bool
	MoveExchange    bool
	MoveObjectOne   bool
	MoveObjectTwo   bool
	MoveAssassinate bool
	MoveSteal       bool
	BlockAmbassador bool
	BlockCaptain    bool
	BlockSubjectOne bool
	BlockSubjectTwo bool
}

func NewState(self *Self, others []*Other, discard *Hand) *State {
	return &State{
		Self:    self,
		Others:  others,
		Discard: discard,
	}
}

type Self struct {
	Coins int
	Hand  *Hand
}

func NewSelf(player *Player) *Self {
	return &Self{
		Coins: player.Coins,
		Hand:  player.Hand,
	}
}

type Other struct {
	Coins int
	Cards int
}

func NewOther(player *Player) *Other {
	return &Other{
		Coins: player.Coins,
		Cards: player.Hand.Size(),
	}
}

type Hand struct {
	Dukes       int
	Assassins   int
	Ambassadors int
	Captains    int
	Contessas   int
}

func NewHand(dukes int, assassins int, ambassadors int, captains int, contessas int) *Hand {
	return &Hand{
		Dukes:       dukes,
		Assassins:   assassins,
		Ambassadors: ambassadors,
		Captains:    captains,
		Contessas:   contessas,
	}
}

// func (h *Hand) Add(card Card) {
//
// }

func (h *Hand) Size() int {
	return h.Dukes + h.Assassins + h.Ambassadors + h.Captains + h.Contessas
}

// type State struct {
// 	Previous *State
// 	Deck     *Deck
// 	Discard  *Pile
// 	Players  []*Player
// }
//
// func NewState(prv *State, deck *Deck, discard *Pile, players []*Player) *State {
// 	return &State{
// 		Previous: prv,
// 		Deck:     deck,
// 		Discard:  discard,
// 		Players:  players,
// 	}
// }
//
// func (s *State) Copy() *State {
//
// 	players := make([]*Player, len(s.Players))
// 	for i, player := range s.Players {
// 		players[i] = player.Copy()
// 	}
//
// 	return NewState(
// 		s,
// 		s.Deck.Copy(),
// 		s.Discard.Copy(),
// 		players,
// 	)
// }
//
// func (s *State) Alive() []*Player {
// 	players := []*Player{}
// 	for _, player := range s.Players {
// 		if player.Alive() {
// 			players = append(players, player)
// 		}
// 	}
// 	return players
// }
