package coup

type Block struct {
	Claim *Claim
}

func NewBlock(sub *Player, dec Type) *Block {
	return &Block{
		Claim: NewClaim(sub, dec, nil),
	}
}

func (b *Block) Successful() bool {
	if b.Claim.Challenge == nil {
		return true
	}

	return b.Claim.Declared == b.Claim.Challenge.Revealed
}

func (b *Block) Scrutinize(state *State) {
	for _, other := range state.Alive()[1:] {
		other.Dispute(b.Claim)
		if b.Claim.Challenge != nil {
			b.Claim.Verify(state)
			return
		}
	}
}
