package coup

type Block struct {
	SubjectOne bool
	SubjectTwo bool
	Ambassador bool
	Captain    bool
}

func NewBlock(sub int, ambassador bool, captain bool) *Block {
	var one bool
	var two bool
	switch sub {
	case 0:
		one = true
	case 1:
		two = true
	}

	return &Block{
		SubjectOne: one,
		SubjectTwo: two,
		Ambassador: ambassador,
		Captain:    captain,
	}
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
