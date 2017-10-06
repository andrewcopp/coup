package coup

type Move struct {
	SubjectOne  bool
	SubjectTwo  bool
	ObjectOne   bool
	ObjectTwo   bool
	Tax         bool
	Exchange    bool
	Assassinate bool
	Steal       bool
}

func NewMove(sub int, obj int, tax bool, exchange bool, assassinate bool, steal bool) *Move {
	var subOne bool
	var subTwo bool
	switch sub {
	case 0:
		subOne = true
	case 1:
		subTwo = true
	}

	var objOne bool
	var objTwo bool
	switch obj {
	case 0:
		objOne = true
	case 1:
		objTwo = true
	}

	return &Move{
		SubjectOne:  subOne,
		SubjectTwo:  subTwo,
		ObjectOne:   objOne,
		ObjectTwo:   objTwo,
		Tax:         tax,
		Exchange:    exchange,
		Assassinate: assassinate,
		Steal:       steal,
	}
}

func (m *Move) Copy() *Move {
	return &Move{
		SubjectOne:  m.SubjectOne,
		SubjectTwo:  m.SubjectTwo,
		ObjectOne:   m.ObjectOne,
		ObjectTwo:   m.ObjectTwo,
		Tax:         m.Tax,
		Exchange:    m.Exchange,
		Assassinate: m.Assassinate,
		Steal:       m.Steal,
	}
}

// type MoveType int
//
// const (
// 	Income      MoveType = iota
// 	ForeignAid           = iota
// 	Coup                 = iota
// 	Tax                  = iota
// 	Assassinate          = iota
// 	Exchange             = iota
// 	Steal                = iota
// )
//
// type Move struct {
// 	Type         MoveType
// 	Subject      *Player
// 	Announcement string
// 	Cost         int
// 	// Claim        *Claim
// 	Counters []CardType
// 	Resolve  func()
// }
//
// func NewMove(t MoveType, sub *Player, announcement string, cost int, claim *Claim, counters []CardType, resolve func()) *Move {
// 	return &Move{
// 		Type:         t,
// 		Subject:      sub,
// 		Announcement: announcement,
// 		Cost:         cost,
// 		// Claim:        claim,
// 		Counters: counters,
// 		Resolve:  resolve,
// 	}
// }
//
// func (m *Move) Successful() bool {
// 	if m.Claim == nil {
// 		return true
// 	}
//
// 	if m.Claim.Challenge == nil {
// 		return true
// 	}
//
// 	return m.Claim.Declared == m.Claim.Challenge.Revealed
// }
//
// func (m *Move) Block(state *State) *Block {
//
// 	var block *Block
//
// 	switch m.Type {
// 	case ForeignAid:
// 		for _, opp := range m.Subject.Opponents(state) {
// 			if block = opp.Impede(m.Counters); block != nil {
// 				break
// 			}
// 		}
// 	case Assassinate, Steal:
// 		block = m.Claim.Object.Impede(m.Counters)
// 	}
//
// 	return block
// }
