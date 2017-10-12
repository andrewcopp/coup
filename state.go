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
