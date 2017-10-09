package coup

type Block struct {
	Subject *Player
	Claim   *Claim
}

func NewBlock(sub *Player, claim *Claim) *Block {
	return &Block{
		Subject: sub,
		Claim:   claim,
	}
}
