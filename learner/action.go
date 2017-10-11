package learner

type Action struct {
	Move           *Move
	ChallengeMove  *ChallengeMove
	Block          *Block
	ChallengeBlock *ChallengeBlock
}

type Move struct {
	Income                                 bool
	ForeignAid                             bool
	Coup                                   bool
	TaxChallengerOneRiskDuke               bool
	TaxChallengerOneRiskAssassin           bool
	TaxChallengerOneRiskAmbassador         bool
	TaxChallengerOneRiskCaptain            bool
	TaxChallengerOneRiskContessa           bool
	TaxChallengerTwoRiskDuke               bool
	TaxChallengerTwoRiskAssassin           bool
	TaxChallengerTwoRiskAmbassador         bool
	TaxChallengerTwoRiskCaptain            bool
	TaxChallengerTwoRiskContessa           bool
	AssassinateChallengerOneRiskDuke       bool
	AssassinateChallengerOneRiskAssassin   bool
	AssassinateChallengerOneRiskAmbassador bool
	AssassinateChallengerOneRiskCaptain    bool
	AssassinateChallengerOneRiskContessa   bool
	AssassinateChallengerTwoRiskDuke       bool
	AssassinateChallengerTwoRiskAssassin   bool
	AssassinateChallengerTwoRiskAmbassador bool
	AssassinateChallengerTwoRiskCaptain    bool
	AssassinateChallengerTwoRiskContessa   bool
	ExchangeChallengerOneRiskDuke          bool
	ExchangeChallengerOneRiskAssassin      bool
	ExchangeChallengerOneRiskAmbassador    bool
	ExchangeChallengerOneRiskCaptain       bool
	ExchangeChallengerOneRiskExchange      bool
	ExchangeChallengerTwoRiskDuke          bool
	ExchangeChallengerTwoRiskAssassin      bool
	ExchangeChallengerTwoRiskAmbassador    bool
	ExchangeChallengerTwoRiskCaptain       bool
	ExchangeChallengerTwoRiskExchange      bool
	StealChallengerOneRiskDuke             bool
	StealChallengerOneRiskAssassin         bool
	StealChallengerOneRiskAmbassador       bool
	StealChallengerOneRiskCaptain          bool
	StealChallengerOneRiskContessa         bool
	StealChallengerTwoRiskDuke             bool
	StealChallengerTwoRiskAssassin         bool
	StealChallengerTwoRiskAmbassador       bool
	StealChallengerTwoRiskCaptain          bool
	StealChallengerTwoRiskContessa         bool
	ObjectOne                              bool
	ObjectTwo                              bool
}

type Block struct {
	RiskDuke       bool
	RiskAssassin   bool
	RiskAmbassador bool
	RiskCaptain    bool
	RiskContessa   bool
	Ambassador     bool
	Captain        bool
}

type ChallengeMove struct {
	SubjectOne                         bool
	SubjectTwo                         bool
	TaxRiskDuke                        bool
	TaxRiskAssassin                    bool
	TaxRiskAmbassador                  bool
	TaxRiskCaptain                     bool
	TaxRiskContessa                    bool
	AssassinateObjectOneRiskDuke       bool
	AssassinateObjectOneRiskAssassin   bool
	AssassinateObjectOneRiskAmbassador bool
	AssassinateObjectOneRiskCaptain    bool
	AssassinateObjectOneRiskContessa   bool
	AssassinateObjectTwoRiskDuke       bool
	AssassinateObjectTwoRiskAssassin   bool
	AssassinateObjectTwoRiskAmbassador bool
	AssassinateObjectTwoRiskCaptain    bool
	AssassinateObjectTwoRiskContessa   bool
	ExchangeRiskDuke                   bool
	ExchangeRiskAssassin               bool
	ExchangeRiskAmbassador             bool
	ExchangeRiskCaptain                bool
	ExchangeRiskContessaz              bool
	StealObjectOneRiskDuke             bool
	StealObjectOneRiskAssassin         bool
	StealObjectOneRiskAmbassador       bool
	StealObjectOneRiskCaptain          bool
	StealObjectOneRiskContessa         bool
	StealObjectTwoRiskDuke             bool
	StealObjectTwoRiskAssassin         bool
	StealObjectTwoRiskAmbassador       bool
	StealObjectTwoRiskCaptain          bool
	StealObjectTwoRiskContessa         bool
}

type ChallengeBlock struct {
	SubjectOneRiskDuke       bool
	SubjectOneRiskAssassin   bool
	SubjectOneRiskAmbassador bool
	SubjectOneRiskCaptain    bool
	SubjectOneRiskContessa   bool
	SubjectTwoRiskDuke       bool
	SubjectTwoRiskAssassin   bool
	SubjectTwoRiskAmbassador bool
	SubjectTwoRiskCaptain    bool
	SubjectTwoRiskContessa   bool
	// TODO: These could be different
	Ambassador bool
	Captain    bool
}

type Discard struct {
	OneDuke        bool
	TwoDukes       bool
	OneAssassin    bool
	TwoAssassins   bool
	OneAmbassador  bool
	TwoAmbassadors bool
	OneCaptain     bool
	TwoCaptains    bool
	OneContessa    bool
	TwoContessas   bool
}
