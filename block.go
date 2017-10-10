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

func (b *Block) Exposed(gm *Game) bool {
	exposed := false
	var challenge *Challenge
	for _, other := range b.Subject.Opponents(gm) {
		if challenge = other.Challenge(gm, b.Claim); challenge != nil {
			if b.Claim.Verify() {
				challenge.Subject.Discard(gm, 1)
				b.Subject.Hand.Remove(b.Claim.Declared)
				card := gm.Deck.Peek()
				gm.Deck.Remove(card)
				b.Subject.Hand.Add(card)
			} else {
				b.Subject.Discard(gm, 1)
				exposed = true
			}
			break
		}
	}

	return exposed
}
