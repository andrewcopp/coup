package coup

type Action struct {
	Move           *ActionMove
	ChallengeMove  *ActionMoveChallenge
	Block          *ActionBlock
	ChallengeBlock *ActionBlockChallenge
}

type Reveal struct {
	Duke       bool
	Assassin   bool
	Ambassador bool
	Captain    bool
	Contessa   bool
}

type Challengeable struct {
	One *Reveal
	Two *Reveal
}

type ActionMove struct {
	Income      bool
	ForeignAid  bool
	Coup        bool
	Tax         *Challengeable
	Assassinate *Challengeable
	Exchange    *Challengeable
	Steal       *Challengeable
	ObjectOne   bool
	ObjectTwo   bool
}

type ActionBlock struct {
	Challengeable *Challengeable
	Ambassador    bool
	Captain       bool
}

type ActionMoveChallenge struct {
	SubjectOne           bool
	SubjectTwo           bool
	Tax                  *Reveal
	AssassinateObjectOne *Reveal
	AssassinateObjectTwo *Reveal
	Exchange             *Reveal
	StealObjectOne       *Reveal
	StealObjectTwo       *Reveal
}

type ActionBlockChallenge struct {
	SubjectOneDuke *Reveal
	SubjectTwoDuke *Reveal
	Ambassador     *Reveal
	Captain        *Reveal
	Contessa       *Reveal
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
