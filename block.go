package coup

// type Block struct {
// 	SubjectOne bool
// 	SubjectTwo bool
// 	Ambassador bool
// 	Captain    bool
// 	Challenge  *Challenge
// }
//
// func NewBlock(sub int, ambassador bool, captain bool, challenge *Challenge) *Block {
// 	var one bool
// 	var two bool
// 	switch sub {
// 	case 0:
// 		one = true
// 	case 1:
// 		two = true
// 	}
//
// 	return &Block{
// 		SubjectOne: one,
// 		SubjectTwo: two,
// 		Ambassador: ambassador,
// 		Captain:    captain,
// 		Challenge:  challenge,
// 	}
// }
//
// func (b *Block) Copy() *Block {
// 	return &Block{
// 		SubjectOne: b.SubjectOne,
// 		SubjectTwo: b.SubjectTwo,
// 		Ambassador: b.Ambassador,
// 		Captain:    b.Captain,
// 	}
// }

func NewBlocks(state *State) []*Action {
	return []*Action{}
}

// type Block struct {
// 	// Claim *Claim
// }
//
// func NewBlock(sub *Player, dec CardType) *Block {
// 	return &Block{
// 	// Claim: NewClaim(sub, dec, nil),
// 	}
// }

// func (b *Block) Successful() bool {
// 	if b.Claim.Challenge == nil {
// 		return true
// 	}
//
// 	return b.Claim.Declared == b.Claim.Challenge.Revealed
// }
