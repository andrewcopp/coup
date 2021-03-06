package coup

// This file represents the complicated state space of the action.
// All helper functions are here as well.

type Action struct {
	Discard        *Discard
	Move           *ActionMove
	ChallengeMove  *ActionMoveChallenge
	Block          *ActionBlock
	ChallengeBlock *ActionBlockChallenge
}

func NewAction() *Action {
	return &Action{
		Discard:        NewDiscard(),
		Move:           NewActionMove(),
		ChallengeMove:  NewActionMoveChallenge(),
		Block:          NewActionBlock(),
		ChallengeBlock: NewActionBlockChallenge(),
	}
}

func (a *Action) Copy() *Action {
	return &Action{
		Discard:        a.Discard.Copy(),
		Move:           a.Move.Copy(),
		ChallengeMove:  a.ChallengeMove.Copy(),
		Block:          a.Block.Copy(),
		ChallengeBlock: a.ChallengeBlock.Copy(),
	}
}

func (a *Action) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, a.Move.Tensor()...)
	tensor = append(tensor, a.ChallengeMove.Tensor()...)
	tensor = append(tensor, a.Block.Tensor()...)
	tensor = append(tensor, a.ChallengeBlock.Tensor()...)
	tensor = append(tensor, a.Discard.Tensor()...)
	return tensor
}

type Reveal struct {
	Duke       bool
	Assassin   bool
	Ambassador bool
	Captain    bool
	Contessa   bool
}

func boolToFloat(prop bool) float64 {
	if prop {
		return 1.0
	}

	return 0.0
}

func (r *Reveal) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, boolToFloat(r.Duke))
	tensor = append(tensor, boolToFloat(r.Assassin))
	tensor = append(tensor, boolToFloat(r.Ambassador))
	tensor = append(tensor, boolToFloat(r.Captain))
	tensor = append(tensor, boolToFloat(r.Contessa))
	return tensor
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

func NewReveals(hand *Cards) []*Reveal {
	reveals := []*Reveal{}
	if hand.Dukes > 0 {
		duke := CardEnum(Duke)
		reveals = append(reveals, NewReveal(&duke))
	}
	if hand.Assassins > 0 {
		assassin := CardEnum(Assassin)
		reveals = append(reveals, NewReveal(&assassin))
	}
	if hand.Ambassadors > 0 {
		ambassador := CardEnum(Ambassador)
		reveals = append(reveals, NewReveal(&ambassador))
	}
	if hand.Captains > 0 {
		captain := CardEnum(Captain)
		reveals = append(reveals, NewReveal(&captain))
	}
	if hand.Contessas > 0 {
		contessa := CardEnum(Contessa)
		reveals = append(reveals, NewReveal(&contessa))
	}
	return reveals
}

func (r *Reveal) Selected() bool {
	if r.Duke || r.Assassin || r.Ambassador || r.Captain || r.Contessa {
		return true
	}

	return false
}

type Challengeable struct {
	One   *Reveal
	Two   *Reveal
	Three *Reveal
	Four  *Reveal
	Five  *Reveal
}

func (c *Challengeable) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, c.One.Tensor()...)
	tensor = append(tensor, c.Two.Tensor()...)
	tensor = append(tensor, c.Three.Tensor()...)
	tensor = append(tensor, c.Four.Tensor()...)
	tensor = append(tensor, c.Five.Tensor()...)
	return tensor
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

func NewChallengeables(subs []int, hand *Cards) []*Challengeable {
	challengeables := []*Challengeable{NewChallengeable(nil, nil, nil, nil, nil)}
	for _, sub := range subs {
		switch sub {
		case 0:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals(hand) {
					c := challengeable.Copy()
					c.One = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 1:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals(hand) {
					c := challengeable.Copy()
					c.Two = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 2:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals(hand) {
					c := challengeable.Copy()
					c.Three = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 3:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals(hand) {
					c := challengeable.Copy()
					c.Four = reveal
					temp = append(temp, c)
				}
			}
			challengeables = temp
		case 4:
			temp := []*Challengeable{}
			for _, challengeable := range challengeables {
				for _, reveal := range NewReveals(hand) {
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

func (c *Challengeable) Selected() bool {
	if c.One.Selected() || c.Two.Selected() || c.Three.Selected() || c.Four.Selected() || c.Five.Selected() {
		return true
	}

	return false
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

func (m *ActionMove) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, boolToFloat(m.Income))
	tensor = append(tensor, boolToFloat(m.ForeignAid))
	tensor = append(tensor, boolToFloat(m.Coup))
	tensor = append(tensor, m.Tax.Tensor()...)
	tensor = append(tensor, m.Assassinate.Tensor()...)
	tensor = append(tensor, m.Exchange.Tensor()...)
	tensor = append(tensor, m.Steal.Tensor()...)
	tensor = append(tensor, boolToFloat(m.ObjectOne))
	tensor = append(tensor, boolToFloat(m.ObjectTwo))
	tensor = append(tensor, boolToFloat(m.ObjectThree))
	tensor = append(tensor, boolToFloat(m.ObjectFour))
	tensor = append(tensor, boolToFloat(m.ObjectFive))
	return tensor
}

func (m *ActionMove) Copy() *ActionMove {
	return &ActionMove{
		Income:      m.Income,
		ForeignAid:  m.ForeignAid,
		Coup:        m.Coup,
		Tax:         m.Tax.Copy(),
		Assassinate: m.Assassinate.Copy(),
		Exchange:    m.Exchange.Copy(),
		Steal:       m.Steal.Copy(),
		ObjectOne:   m.ObjectOne,
		ObjectTwo:   m.ObjectTwo,
		ObjectThree: m.ObjectThree,
		ObjectFour:  m.ObjectFour,
		ObjectFive:  m.ObjectFive,
	}
}

func NewActionMove() *ActionMove {
	return &ActionMove{
		Tax:         NewChallengeable(nil, nil, nil, nil, nil),
		Assassinate: NewChallengeable(nil, nil, nil, nil, nil),
		Exchange:    NewChallengeable(nil, nil, nil, nil, nil),
		Steal:       NewChallengeable(nil, nil, nil, nil, nil),
	}
}

type ActionBlock struct {
	Challengeable *Challengeable
	Ambassador    bool
	Captain       bool
}

func (b *ActionBlock) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, b.Challengeable.Tensor()...)
	tensor = append(tensor, boolToFloat(b.Ambassador))
	tensor = append(tensor, boolToFloat(b.Captain))
	return tensor
}

func NewActionBlock() *ActionBlock {
	return &ActionBlock{
		Challengeable: NewChallengeable(
			NewReveal(nil),
			NewReveal(nil),
			NewReveal(nil),
			NewReveal(nil),
			NewReveal(nil),
		),
	}
}

func (b *ActionBlock) Copy() *ActionBlock {
	return &ActionBlock{
		Challengeable: b.Challengeable.Copy(),
		Ambassador:    b.Ambassador,
		Captain:       b.Captain,
	}
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

func (c *ActionMoveChallenge) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, boolToFloat(c.SubjectOne))
	tensor = append(tensor, boolToFloat(c.SubjectTwo))
	tensor = append(tensor, boolToFloat(c.SubjectThree))
	tensor = append(tensor, boolToFloat(c.SubjectFour))
	tensor = append(tensor, boolToFloat(c.SubjectFive))
	tensor = append(tensor, c.Tax.Tensor()...)
	tensor = append(tensor, c.AssassinateObjectOne.Tensor()...)
	tensor = append(tensor, c.AssassinateObjectTwo.Tensor()...)
	tensor = append(tensor, c.AssassinateObjectThree.Tensor()...)
	tensor = append(tensor, c.AssassinateObjectFour.Tensor()...)
	tensor = append(tensor, c.AssassinateObjectFive.Tensor()...)
	tensor = append(tensor, c.Exchange.Tensor()...)
	tensor = append(tensor, c.StealObjectOne.Tensor()...)
	tensor = append(tensor, c.StealObjectTwo.Tensor()...)
	tensor = append(tensor, c.StealObjectThree.Tensor()...)
	tensor = append(tensor, c.StealObjectFour.Tensor()...)
	tensor = append(tensor, c.StealObjectFive.Tensor()...)
	return tensor
}

func NewActionMoveChallenge() *ActionMoveChallenge {
	return &ActionMoveChallenge{
		Tax:                    NewReveal(nil),
		AssassinateObjectOne:   NewReveal(nil),
		AssassinateObjectTwo:   NewReveal(nil),
		AssassinateObjectThree: NewReveal(nil),
		AssassinateObjectFour:  NewReveal(nil),
		AssassinateObjectFive:  NewReveal(nil),
		Exchange:               NewReveal(nil),
		StealObjectOne:         NewReveal(nil),
		StealObjectTwo:         NewReveal(nil),
		StealObjectThree:       NewReveal(nil),
		StealObjectFour:        NewReveal(nil),
		StealObjectFive:        NewReveal(nil),
	}
}

func (c *ActionMoveChallenge) Copy() *ActionMoveChallenge {
	return &ActionMoveChallenge{
		SubjectOne:             c.SubjectOne,
		SubjectTwo:             c.SubjectTwo,
		SubjectThree:           c.SubjectThree,
		SubjectFour:            c.SubjectFour,
		SubjectFive:            c.SubjectFive,
		Tax:                    c.Tax.Copy(),
		AssassinateObjectOne:   c.AssassinateObjectOne.Copy(),
		AssassinateObjectTwo:   c.AssassinateObjectTwo.Copy(),
		AssassinateObjectThree: c.AssassinateObjectThree.Copy(),
		AssassinateObjectFour:  c.AssassinateObjectFour.Copy(),
		AssassinateObjectFive:  c.AssassinateObjectFive.Copy(),
		Exchange:               c.Exchange.Copy(),
		StealObjectOne:         c.StealObjectOne.Copy(),
		StealObjectTwo:         c.StealObjectTwo.Copy(),
		StealObjectThree:       c.StealObjectThree.Copy(),
		StealObjectFour:        c.StealObjectFour.Copy(),
		StealObjectFive:        c.StealObjectFive.Copy(),
	}
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

func (c *ActionBlockChallenge) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, c.SubjectOneDuke.Tensor()...)
	tensor = append(tensor, c.SubjectTwoDuke.Tensor()...)
	tensor = append(tensor, c.SubjectThreeDuke.Tensor()...)
	tensor = append(tensor, c.SubjectFourDuke.Tensor()...)
	tensor = append(tensor, c.SubjectFiveDuke.Tensor()...)
	tensor = append(tensor, c.Ambassador.Tensor()...)
	tensor = append(tensor, c.Captain.Tensor()...)
	tensor = append(tensor, c.Contessa.Tensor()...)
	return tensor
}

func NewActionBlockChallenge() *ActionBlockChallenge {
	return &ActionBlockChallenge{
		SubjectOneDuke:   NewReveal(nil),
		SubjectTwoDuke:   NewReveal(nil),
		SubjectThreeDuke: NewReveal(nil),
		SubjectFourDuke:  NewReveal(nil),
		SubjectFiveDuke:  NewReveal(nil),
		Ambassador:       NewReveal(nil),
		Captain:          NewReveal(nil),
		Contessa:         NewReveal(nil),
	}
}

func (c *ActionBlockChallenge) Copy() *ActionBlockChallenge {
	return &ActionBlockChallenge{
		SubjectOneDuke:   c.SubjectOneDuke,
		SubjectTwoDuke:   c.SubjectTwoDuke,
		SubjectThreeDuke: c.SubjectThreeDuke,
		SubjectFourDuke:  c.SubjectFourDuke,
		SubjectFiveDuke:  c.SubjectFiveDuke,
		Ambassador:       c.Ambassador,
		Captain:          c.Captain,
		Contessa:         c.Contessa,
	}
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

func (d *Discard) Tensor() []float64 {
	tensor := []float64{}
	tensor = append(tensor, boolToFloat(d.OneDuke))
	tensor = append(tensor, boolToFloat(d.TwoDukes))
	tensor = append(tensor, boolToFloat(d.OneAssassin))
	tensor = append(tensor, boolToFloat(d.TwoAssassins))
	tensor = append(tensor, boolToFloat(d.OneAmbassador))
	tensor = append(tensor, boolToFloat(d.TwoAmbassadors))
	tensor = append(tensor, boolToFloat(d.OneCaptain))
	tensor = append(tensor, boolToFloat(d.TwoCaptains))
	tensor = append(tensor, boolToFloat(d.OneContessa))
	tensor = append(tensor, boolToFloat(d.TwoContessas))
	return tensor
}

func NewDiscard() *Discard {
	return &Discard{}
}

func (d *Discard) Copy() *Discard {
	return &Discard{
		OneDuke:        d.OneDuke,
		TwoDukes:       d.TwoDukes,
		OneAssassin:    d.OneAssassin,
		TwoAssassins:   d.TwoAssassins,
		OneAmbassador:  d.OneAmbassador,
		TwoAmbassadors: d.TwoAmbassadors,
		OneCaptain:     d.OneCaptain,
		TwoCaptains:    d.TwoCaptains,
		OneContessa:    d.OneContessa,
		TwoContessas:   d.TwoContessas,
	}
}
