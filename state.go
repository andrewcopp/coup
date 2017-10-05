package coup

type State struct {
	Self    *Self
	Others  []*Other
	Discard *Hand
	Move    *Move
	Block   *Block
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

// func (s *State) Alive() []*Player {
// 	players := []*Player{}
// 	for _, player := range s.Players {
// 		if player.Alive() {
// 			players = append(players, player)
// 		}
// 	}
// 	return players
// }
