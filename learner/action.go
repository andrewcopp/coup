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
}
