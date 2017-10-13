package coup

import "fmt"

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
		if b.Challenge = other.ChallengeBlock(gm, b.Claim, nil); b.Challenge != nil {
			if gm.Logs {
				fmt.Printf("%s challenges.\n", b.Challenge.Subject.Name)
			}
			if b.Claim.Verify() {
				if gm.Logs {
					fmt.Printf("Challenge unsuccessful.\n")
				}
				for _, card := range b.Challenge.Subject.Discard(gm, 1) {
					gm.Discard.Add(card)
				}
				b.Subject.Hand.Remove(b.Claim.Declared)
				gm.Deck.Add(b.Claim.Declared)
				card := gm.Deck.Peek()
				gm.Deck.Remove(card)
				b.Subject.Hand.Add(card)
			} else {
				if gm.Logs {
					fmt.Printf("Challenge successful.\n")
				}
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
