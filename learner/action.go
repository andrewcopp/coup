package learner

type Action struct {
	MoveIncome                    bool
	MoveForeignAid                bool
	MoveCoup                      bool
	MoveTax                       bool
	MoveAssassinate               bool
	MoveExchange                  bool
	MoveSteal                     bool
	MoveObjectOne                 bool
	MoveObjectTwo                 bool
	ChallengeSubjectOne           bool
	ChallengeSubjectTwo           bool
	ChallengeTax                  bool
	ChallengeAssassinateObjectOne bool
	ChallengeAssassinateObjectTwo bool
	ChallengeExchange             bool
	ChallengeStealObjectOne       bool
	ChallengeStealObjectTwo       bool
	Block                         bool
	BlockAmbassador               bool
	BlockCaptain                  bool
	ChallengeBlockSubjectOne      bool
	ChallengeBlockSubjectTwo      bool
	ChallengeBlockAmbassador      bool
	ChallengeBlockCaptain         bool
	DiscardOneDuke                bool
	DiscardTwoDukes               bool
	DiscardOneAssassin            bool
	DiscardTwoAssassins           bool
	DiscardOneAmbassador          bool
	DiscardTwoAmbassadors         bool
	DiscardOneCaptain             bool
	DiscardTwoCaptains            bool
	DiscardOneContessa            bool
	DiscardTwoContessas           bool
}
