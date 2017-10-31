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
	ThreeCoins           float64
	ThreeCards           float64
	FourCoins            float64
	FourCards            float64
	FiveCoins            float64
	FiveCards            float64
	DiscardedDukes       float64
	DiscardedAssassins   float64
	DiscardedAmbassadors float64
	DiscardedCaptains    float64
	DiscardedContessas   float64
	Bridge               *Bridge
}

type Bridge struct {
	SubjectOne                 bool
	SubjectTwo                 bool
	SubjectThree               bool
	SubjectFour                bool
	SubjectFive                bool
	MoveForeignAid             bool
	MoveTax                    bool
	MoveAssassinate            bool
	MoveExchange               bool
	MoveSteal                  bool
	ObjectOne                  bool
	ObjectTwo                  bool
	ObjectThree                bool
	ObjectFour                 bool
	ObjectFive                 bool
	MoveChallengeSubjectOne    bool
	MoveChallengeSubjectTwo    bool
	MoveChallengeSubjectThree  bool
	MoveChallengeSubjectFour   bool
	MoveChallengeSubjectFive   bool
	BlockSubjectOne            bool
	BlockSubjectTwo            bool
	BlockSubjectThree          bool
	BlockSubjectFour           bool
	BlockSubjectFive           bool
	BlockAmbassador            bool
	BlockCaptain               bool
	BlockChallengeSubjectOne   bool
	BlockChallengeSubjectTwo   bool
	BlockChallengeSubjectThree bool
	BlockChallengeSubjectFour  bool
	BlockChallengeSubjectFive  bool
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
		ThreeCoins:           float64(players[2].Coins) / 12.0,
		ThreeCards:           float64(players[2].Hand.Size()) / 4.0,
		FourCoins:            float64(players[3].Coins) / 12.0,
		FourCards:            float64(players[3].Hand.Size()) / 4.0,
		FiveCoins:            float64(players[4].Coins) / 12.0,
		FiveCards:            float64(players[4].Hand.Size()) / 4.0,
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
		case players[2]:
			state.Bridge.SubjectThree = true
		case players[3]:
			state.Bridge.SubjectFour = true
		case players[4]:
			state.Bridge.SubjectFive = true
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
			case players[2]:
				state.Bridge.ObjectThree = true
			case players[3]:
				state.Bridge.ObjectFour = true
			case players[4]:
				state.Bridge.ObjectFive = true
			}
		}

		if mv.Challenge != nil {
			switch mv.Challenge.Subject {
			case players[0]:
				state.Bridge.MoveChallengeSubjectOne = true
			case players[1]:
				state.Bridge.MoveChallengeSubjectTwo = true
			case players[2]:
				state.Bridge.MoveChallengeSubjectThree = true
			case players[3]:
				state.Bridge.MoveChallengeSubjectFour = true
			case players[4]:
				state.Bridge.MoveChallengeSubjectFive = true
			}
		}
	}

	if blk != nil {
		switch blk.Subject {
		case players[0]:
			state.Bridge.BlockSubjectOne = true
		case players[1]:
			state.Bridge.BlockSubjectTwo = true
		case players[2]:
			state.Bridge.BlockSubjectThree = true
		case players[3]:
			state.Bridge.BlockSubjectFour = true
		case players[4]:
			state.Bridge.BlockSubjectFive = true
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
			case players[2]:
				state.Bridge.BlockChallengeSubjectThree = true
			case players[3]:
				state.Bridge.BlockChallengeSubjectFour = true
			case players[4]:
				state.Bridge.BlockChallengeSubjectFive = true
			}
		}
	}

	return state
}

func (s *State) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, s.OneCoins)
	tensor = append(tensor, s.OneDukes)
	tensor = append(tensor, s.OneAssassins)
	tensor = append(tensor, s.OneAmbassadors)
	tensor = append(tensor, s.OneCaptains)
	tensor = append(tensor, s.OneContessas)
	tensor = append(tensor, s.TwoCoins)
	tensor = append(tensor, s.TwoCards)
	tensor = append(tensor, s.ThreeCoins)
	tensor = append(tensor, s.ThreeCards)
	tensor = append(tensor, s.FourCoins)
	tensor = append(tensor, s.FourCards)
	tensor = append(tensor, s.FiveCoins)
	tensor = append(tensor, s.FiveCards)
	tensor = append(tensor, s.DiscardedDukes)
	tensor = append(tensor, s.DiscardedAssassins)
	tensor = append(tensor, s.DiscardedAmbassadors)
	tensor = append(tensor, s.DiscardedCaptains)
	tensor = append(tensor, s.DiscardedContessas)
	tensor = append(tensor, s.boolToFloat(s.Bridge.SubjectOne))
	tensor = append(tensor, s.boolToFloat(s.Bridge.SubjectTwo))
	tensor = append(tensor, s.boolToFloat(s.Bridge.SubjectThree))
	tensor = append(tensor, s.boolToFloat(s.Bridge.SubjectFour))
	tensor = append(tensor, s.boolToFloat(s.Bridge.SubjectFive))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveForeignAid))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveTax))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveAssassinate))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveExchange))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveSteal))
	tensor = append(tensor, s.boolToFloat(s.Bridge.ObjectOne))
	tensor = append(tensor, s.boolToFloat(s.Bridge.ObjectTwo))
	tensor = append(tensor, s.boolToFloat(s.Bridge.ObjectThree))
	tensor = append(tensor, s.boolToFloat(s.Bridge.ObjectFour))
	tensor = append(tensor, s.boolToFloat(s.Bridge.ObjectFive))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveChallengeSubjectOne))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveChallengeSubjectTwo))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveChallengeSubjectThree))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveChallengeSubjectFour))
	tensor = append(tensor, s.boolToFloat(s.Bridge.MoveChallengeSubjectFive))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockSubjectOne))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockSubjectTwo))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockSubjectThree))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockSubjectFour))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockSubjectFive))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockAmbassador))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockCaptain))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockChallengeSubjectOne))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockChallengeSubjectTwo))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockChallengeSubjectThree))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockChallengeSubjectFour))
	tensor = append(tensor, s.boolToFloat(s.Bridge.BlockChallengeSubjectFive))
	return tensor

}

func (s *State) boolToFloat(prop bool) float64 {
	if prop {
		return 1.0
	}

	return 0.0
}
