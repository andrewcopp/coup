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

func NewReveal(card *CardEnum) *Reveal {
	reveal := &Reveal{}
	if card != nil {
		switch *card {
		case Duke:
			reveal.Duke = true
		case Assassin:
			reveal.Assassin = true
		case Ambassador:
			reveal.Ambassador = true
		case Captain:
			reveal.Captain = true
		case Contessa:
			reveal.Contessa = true
		}
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
	duke := CardEnum(Duke)
	assassin := CardEnum(Assassin)
	ambassador := CardEnum(Ambassador)
	captain := CardEnum(Captain)
	contessa := CardEnum(Contessa)
	return []*Reveal{
		NewReveal(&duke),
		NewReveal(&assassin),
		NewReveal(&ambassador),
		NewReveal(&captain),
		NewReveal(&contessa),
	}
}

type Challengeable struct {
	One   *Reveal
	Two   *Reveal
	Three *Reveal
	Four  *Reveal
	Five  *Reveal
}

func NewChallengeable(one *Reveal, two *Reveal, three *Reveal, four *Reveal, five *Reveal) *Challengeable {
	if one == nil {
		one = NewReveal(nil)
	}

	if two == nil {
		two = NewReveal(nil)
	}

	if three == nil {
		three = NewReveal(nil)
	}

	if four == nil {
		four = NewReveal(nil)
	}

	if five == nil {
		five = NewReveal(nil)
	}

	return &Challengeable{
		One:   one,
		Two:   two,
		Three: three,
		Four:  four,
		Five:  five,
	}
}

func (c *Challengeable) Copy() *Challengeable {
	var one *Reveal
	if c.One != nil {
		one = c.One.Copy()
	}

	var two *Reveal
	if c.Two != nil {
		two = c.Two.Copy()
	}

	var three *Reveal
	if c.Three != nil {
		three = c.Three.Copy()
	}

	var four *Reveal
	if c.Four != nil {
		four = c.Four.Copy()
	}

	var five *Reveal
	if c.Five != nil {
		five = c.Five.Copy()
	}

	return &Challengeable{
		One:   one,
		Two:   two,
		Three: three,
		Four:  four,
		Five:  five,
	}
}

func NewChallengeables(subs []int) []*Challengeable {
	challengeables := []*Challengeable{NewChallengeable(nil, nil, nil, nil, nil)}
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
		case 2:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals() {
					c := challengeable.Copy()
					c.Three = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 3:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals() {
					c := challengeable.Copy()
					c.Four = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 4:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals() {
					c := challengeable.Copy()
					c.Five = reveal
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
	ObjectThree bool
	ObjectFour  bool
	ObjectFive  bool
}

type ActionBlock struct {
	Challengeable *Challengeable
	Ambassador    bool
	Captain       bool
}

type ActionMoveChallenge struct {
	SubjectOne             bool
	SubjectTwo             bool
	SubjectThree           bool
	SubjectFour            bool
	SubjectFive            bool
	Tax                    *Reveal
	AssassinateObjectOne   *Reveal
	AssassinateObjectTwo   *Reveal
	AssassinateObjectThree *Reveal
	AssassinateObjectFour  *Reveal
	AssassinateObjectFive  *Reveal
	Exchange               *Reveal
	StealObjectOne         *Reveal
	StealObjectTwo         *Reveal
	StealObjectThree       *Reveal
	StealObjectFour        *Reveal
	StealObjectFive        *Reveal
}

type ActionBlockChallenge struct {
	SubjectOneDuke   *Reveal
	SubjectTwoDuke   *Reveal
	SubjectThreeDuke *Reveal
	SubjectFourDuke  *Reveal
	SubjectFiveDuke  *Reveal
	Ambassador       *Reveal
	Captain          *Reveal
	Contessa         *Reveal
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
