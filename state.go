package coup

type State struct {
	OneCoins             float64
	OneDukes             float64
	OneAssassins         float64
	OneAmbassadors       float64
	OneCaptains          float64
	OneContessas         float64
	TwoCoins             float64
	TwoCards             float64
	DiscardedDukes       float64
	DiscardedAssassins   float64
	DiscardedAmbassadors float64
	DiscardedCaptains    float64
	DiscardedContessas   float64
	Bridge               *Bridge
}

type Bridge struct {
	SubjectOne               bool
	SubjectTwo               bool
	MoveForeignAid           bool
	MoveTax                  bool
	MoveAssassinate          bool
	MoveExchange             bool
	MoveSteal                bool
	ObjectOne                bool
	ObjectTwo                bool
	MoveChallengeSubjectOne  bool
	MoveChallengeSubjectTwo  bool
	BlockSubjectOne          bool
	BlockSubjectTwo          bool
	BlockAmbassador          bool
	BlockCaptain             bool
	BlockChallengeSubjectOne bool
	BlockChallengeSubjectTwo bool
}

func NewState(self *Player, gm *Game, mv *Move, blk *Block) *State {
	index := 0
	for i, player := range gm.Players {
		if self == player {
			index = i
		}
	}

	players := append(gm.Players[index:], gm.Players[:index]...)
	state := &State{
		OneCoins:             float64(players[0].Coins) / 12.0,
		OneDukes:             float64(players[0].Hand.Dukes) / 3.0,
		OneAssassins:         float64(players[0].Hand.Assassins) / 3.0,
		OneAmbassadors:       float64(players[0].Hand.Ambassadors) / 3.0,
		OneCaptains:          float64(players[0].Hand.Captains) / 3.0,
		OneContessas:         float64(players[0].Hand.Contessas) / 3.0,
		TwoCoins:             float64(players[1].Coins) / 12.0,
		TwoCards:             float64(players[1].Hand.Size()) / 4.0,
		DiscardedDukes:       float64(gm.Discard.Dukes) / 3.0,
		DiscardedAssassins:   float64(gm.Discard.Assassins) / 3.0,
		DiscardedAmbassadors: float64(gm.Discard.Ambassadors) / 3.0,
		DiscardedCaptains:    float64(gm.Discard.Captains) / 3.0,
		DiscardedContessas:   float64(gm.Discard.Contessas) / 3.0,
		Bridge:               &Bridge{},
	}

	if mv != nil {
		switch mv.Subject {
		case players[0]:
			state.Bridge.SubjectOne = true
		case players[1]:
			state.Bridge.SubjectTwo = true
		}

		switch mv.Case {
		case ForeignAid:
			state.Bridge.MoveForeignAid = true
		case Tax:
			state.Bridge.MoveTax = true
		case Assassinate:
			state.Bridge.MoveAssassinate = true
		case Exchange:
			state.Bridge.MoveExchange = true
		case Steal:
			state.Bridge.MoveSteal = true
		}

		if mv.Object != nil {
			switch mv.Object {
			case players[0]:
				state.Bridge.ObjectOne = true
			case players[1]:
				state.Bridge.ObjectTwo = true
			}
		}

		if mv.Challenge != nil {
			switch mv.Challenge.Subject {
			case players[0]:
				state.Bridge.MoveChallengeSubjectOne = true
			case players[1]:
				state.Bridge.MoveChallengeSubjectTwo = true
			}
		}
	}

	if blk != nil {
		switch blk.Subject {
		case players[0]:
			state.Bridge.BlockSubjectOne = true
		case players[1]:
			state.Bridge.BlockSubjectTwo = true
		}

		switch blk.Claim.Declared {
		case Ambassador:
			state.Bridge.BlockAmbassador = true
		case Captain:
			state.Bridge.BlockCaptain = true
		}

		if blk.Challenge != nil {
			switch blk.Challenge.Subject {
			case players[0]:
				state.Bridge.BlockChallengeSubjectOne = true
			case players[1]:
				state.Bridge.BlockChallengeSubjectTwo = true
			}
		}
	}

	return state
}
