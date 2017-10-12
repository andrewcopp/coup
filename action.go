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

func NewReveal(card CardEnum) *Reveal {
	reveal := &Reveal{}
	switch card {
	case Duke:
		reveal.Duke = true
	case Assassin:
		reveal.Assassin = true
	case Ambassador:
		reveal.Assassin = true
	case Captain:
		reveal.Captain = true
	case Contessa:
		reveal.Contessa = true
	}
	return reveal
}

func (r *Reveal) Copy() *Reveal {
	return &Reveal{
		Duke:       r.Duke,
		Assassin:   r.Assassin,
		Ambassador: r.Ambassador,
		Captain:    r.Captain,
		Contessa:   r.Contessa,
	}
}

func NewReveals() []*Reveal {
	return []*Reveal{
		NewReveal(Duke),
		NewReveal(Assassin),
		NewReveal(Ambassador),
		NewReveal(Captain),
		NewReveal(Contessa),
	}
}

type Challengeable struct {
	One *Reveal
	Two *Reveal
}

func NewChallengeable(one *Reveal, two *Reveal) *Challengeable {
	return &Challengeable{
		One: one,
		Two: two,
	}
}

func (c *Challengeable) Copy() *Challengeable {
	return &Challengeable{
		One: c.One.Copy(),
		Two: c.Two.Copy(),
	}
}

func NewChallengeables(subs []int) []*Challengeable {
	challengeables := []*Challengeable{NewChallengeable(nil, nil)}
	for _, sub := range subs {
		switch sub {
		case 0:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals() {
					c := challengeable.Copy()
					c.One = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 1:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals() {
					c := challengeable.Copy()
					c.Two = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		}
	}
	return challengeables
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
