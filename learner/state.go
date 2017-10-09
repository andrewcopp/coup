package learner

type State struct {
	OneCoins             int
	OneDukes             int
	OneAssassins         int
	OneAmbassadors       int
	OneCaptains          int
	OneContessas         int
	TwoCoins             int
	TwoCards             int
	DiscardedDukes       int
	DiscardedAssassins   int
	DiscardedAmbassadors int
	DiscardedCaptains    int
	DiscardedContessas   int
	SubjectOne           bool
	SubjectTwo           bool
	MoveForeignAid       bool
	MoveAssassinate      bool
	MoveSteal            bool
	ObjectOne            bool
	ObjectTwo            bool
}

// type State struct {
// 	Self    *Self
// 	Others  []*Other
// 	Discard *Hand
// 	Turn    *Turn
// }
//
// func NewState(self *Self, others []*Other, discard *Hand) *State {
// 	return &State{
// 		Self:    self,
// 		Others:  others,
// 		Discard: discard,
// 		Turn:    NewTurn(),
// 	}
// }
//
// func (s *State) Copy() *State {
// 	others := make([]*Other, len(s.Others))
// 	for i, other := range s.Others {
// 		others[i] = other.Copy()
// 	}
//
// 	return &State{
// 		Self:    s.Self.Copy(),
// 		Others:  others,
// 		Discard: s.Discard.Copy(),
// 		Turn:    s.Turn.Copy(),
// 	}
// }
//
// type Self struct {
// 	Coins int
// 	Hand  *Hand
// }
//
// func NewSelf(player *Player) *Self {
// 	return &Self{
// 		Coins: player.Coins,
// 		Hand:  player.Hand,
// 	}
// }
//
// func (s *Self) Copy() *Self {
// 	return &Self{
// 		Coins: s.Coins,
// 		Hand:  s.Hand.Copy(),
// 	}
// }
//
// type Other struct {
// 	Coins int
// 	Cards int
// }
//
// func NewOther(player *Player) *Other {
// 	return &Other{
// 		Coins: player.Coins,
// 		Cards: player.Hand.Size(),
// 	}
// }
//
// func (o *Other) Copy() *Other {
// 	return &Other{
// 		Coins: o.Coins,
// 		Cards: o.Cards,
// 	}
// }
//
// type Hand struct {
// 	Dukes       int
// 	Assassins   int
// 	Ambassadors int
// 	Captains    int
// 	Contessas   int
// }
//
// func NewHand(dukes int, assassins int, ambassadors int, captains int, contessas int) *Hand {
// 	return &Hand{
// 		Dukes:       dukes,
// 		Assassins:   assassins,
// 		Ambassadors: ambassadors,
// 		Captains:    captains,
// 		Contessas:   contessas,
// 	}
// }
//
// func (h *Hand) Copy() *Hand {
// 	return &Hand{
// 		Dukes:       h.Dukes,
// 		Assassins:   h.Assassins,
// 		Ambassadors: h.Ambassadors,
// 		Captains:    h.Captains,
// 		Contessas:   h.Contessas,
// 	}
// }
//
// // func (h *Hand) Add(card Card) {
// //
// // }
//
// func (h *Hand) Size() int {
// 	return h.Dukes + h.Assassins + h.Ambassadors + h.Captains + h.Contessas
// }
//
// // func (s *State) Alive() []*Player {
// // 	players := []*Player{}
// // 	for _, player := range s.Players {
// // 		if player.Alive() {
// // 			players = append(players, player)
// // 		}
// // 	}
// // 	return players
// // }
