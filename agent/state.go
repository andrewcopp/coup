package agent

import "github.com/andrewcopp/coup"

type State struct {
	OneCoins             int
	OneDukes             int
	OneAssassins         int
	OneAmbassadors       int
	OneCaptains          int
	OneContessa          int
	TwoCoins             int
	TwoCards             int
	DiscardedDukes       int
	DiscardedAssassins   int
	DiscardedAmbassadors int
	DiscardedCaptains    int
	DiscardedContessas   int
	MoveSubjectOne       bool
	MoveSubjectTwo       bool
	MoveTax              bool
	MoveExchange         bool
	MoveObjectOne        bool
	MoveObjectTwo        bool
	MoveAssassinate      bool
	MoveSteal            bool
	BlockAmbassador      bool
	BlockCaptain         bool
	BlockSubjectOne      bool
	BlockSubjectTwo      bool
}

func NewState(s *coup.State) *State {
	return &State{}
}
