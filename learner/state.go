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
