package coup

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
