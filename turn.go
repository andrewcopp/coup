package coup

type Turn struct {
	MoveSubjectOne           bool
	MoveSubjectTwo           bool
	MoveObjectOne            bool
	MoveObjectTwo            bool
	MoveForeignAid           bool
	MoveTax                  bool
	MoveExchange             bool
	MoveAssassinate          bool
	MoveSteal                bool
	MoveChallengeSubjectOne  bool
	MoveChallengeSubjectTwo  bool
	BlockSubjectOne          bool
	BlockSubjectTwo          bool
	BlockAmbassador          bool
	BlockCaptain             bool
	BlockChallengeSubjectOne bool
	BlockChallengeSubjectTwo bool
}

func NewTurn() *Turn {
	return &Turn{}
}

func (t *Turn) Copy() *Turn {
	return &Turn{
		MoveSubjectOne:           t.MoveSubjectOne,
		MoveSubjectTwo:           t.MoveSubjectTwo,
		MoveObjectOne:            t.MoveObjectOne,
		MoveObjectTwo:            t.MoveObjectTwo,
		MoveForeignAid:           t.MoveForeignAid,
		MoveTax:                  t.MoveTax,
		MoveExchange:             t.MoveExchange,
		MoveAssassinate:          t.MoveAssassinate,
		MoveSteal:                t.MoveSteal,
		MoveChallengeSubjectOne:  t.MoveChallengeSubjectOne,
		MoveChallengeSubjectTwo:  t.MoveChallengeSubjectTwo,
		BlockSubjectOne:          t.BlockSubjectOne,
		BlockSubjectTwo:          t.BlockSubjectTwo,
		BlockAmbassador:          t.BlockAmbassador,
		BlockCaptain:             t.BlockCaptain,
		BlockChallengeSubjectOne: t.BlockChallengeSubjectOne,
		BlockChallengeSubjectTwo: t.BlockChallengeSubjectTwo,
	}
}

func (t *Turn) Move(sub int, obj int, foreignAid bool, tax bool, exchange bool, assassinate bool, steal bool) {
	switch sub {
	case 0:
		t.MoveSubjectOne = true
	case 1:
		t.MoveSubjectTwo = true
	}

	switch obj {
	case 0:
		t.MoveObjectOne = true
	case 1:
		t.MoveObjectTwo = true
	}

	t.MoveForeignAid = foreignAid
	t.MoveTax = tax
	t.MoveExchange = exchange
	t.MoveAssassinate = assassinate
	t.MoveSteal = steal
}

func (t *Turn) Challenge(sub int) {
	if !t.BlockSubjectOne && !t.BlockSubjectTwo {
		switch sub {
		case 0:
			t.MoveChallengeSubjectOne = true
		case 1:
			t.MoveChallengeSubjectTwo = true
		}
	} else {
		switch sub {
		case 0:
			t.BlockChallengeSubjectOne = true
		case 1:
			t.BlockChallengeSubjectTwo = true
		}
	}
}

func (t *Turn) Block(sub int, ambassador bool, captain bool) {
	switch sub {
	case 0:
		t.BlockSubjectOne = true
	case 1:
		t.BlockSubjectTwo = true
	}

	t.BlockAmbassador = ambassador
	t.BlockCaptain = captain
}
