package agent

import "github.com/andrewcopp/coup"

type State struct {
	OneCoins                int
	OneDukes                int
	OneAssassins            int
	OneAmbassadors          int
	OneCaptains             int
	OneContessa             int
	TwoCoins                int
	TwoCards                int
	DiscardedDukes          int
	DiscardedAssassins      int
	DiscardedAmbassadors    int
	DiscardedCaptains       int
	DiscardedContessas      int
	MoveForeignAid          bool
	MoveTax                 bool
	MoveAssassinate         bool
	MoveExchange            bool
	MoveSteal               bool
	MoveSubjectOne          bool
	MoveSubjectTwo          bool
	MoveObjectOne           bool
	MoveObjectTwo           bool
	MoveChallengerNone      bool
	MoveChallengerOne       bool
	MoveChallengerTwo       bool
	MoveChallengeSuccessful bool
	BlockDuke               bool
	BlockAmbassador         bool
	BlockCaptain            bool
	BlockContessa           bool
	BlockSubjectNone        bool
	BlockSubjectOne         bool
	BlockSubjectTwo         bool
	BlockChallengerNone     bool
	BlockChallengerOne      bool
	BLockChallengerTwo      bool
}

func NewState(s *coup.State) *State {
	return &State{}
}
