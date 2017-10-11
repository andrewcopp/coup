package coup

type Block struct {
	Subject   *Player
	Claim     *Claim
	Challenge *Challenge
}

func NewBlock(sub *Player, claim *Claim) *Block {
	return &Block{
		Subject: sub,
		Claim:   claim,
	}
}

func (b *Block) Exposed(gm *Game) bool {
	exposed := false
	for _, other := range b.Subject.Opponents(gm) {
		if b.Challenge = other.Challenge(gm, b.Claim); b.Challenge != nil {
			if b.Claim.Verify() {
				for _, card := range b.Challenge.Subject.Discard(gm, 1) {
					gm.Discard.Add(card)
				}
				b.Subject.Hand.Remove(b.Claim.Declared)
				gm.Deck.Add(b.Claim.Declared)
				card := gm.Deck.Peek()
				gm.Deck.Remove(card)
				b.Subject.Hand.Add(card)
			} else {
				for _, card := range b.Subject.Discard(gm, 1) {
					gm.Discard.Add(card)
				}
				exposed = true
			}
			break
		}
	}

	return exposed
}
