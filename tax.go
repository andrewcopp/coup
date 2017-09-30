package coup

import "fmt"

type Tax struct {
	Subject   *Player
	Challenge *Challenge
}

func NewTax(sub *Player) *Tax {
	return &Tax{
		Subject:   sub,
		Challenge: nil,
	}
}

func (t *Tax) Modify(state *State) {
	fmt.Printf("%s taxes.\n", t.Subject.Name)

	claim := NewClaim(t.Subject, Duke, nil)
	for _, other := range state.Alive()[1:] {
		if claim.Challenger == nil {
			other.Dispute(claim)
		}
	}

	if claim.Challenger != nil {
		fmt.Printf("%s challenges.\n", claim.Challenger.Name)
		if card := t.Subject.Produce(claim.Declared); card != nil {
			fmt.Printf("Challenge unsuccessful.\n")
			claim.Challenger.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", claim.Challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
			state.Deck.Add(card)
			t.Subject.Hand = append(t.Subject.Hand, state.Deck.Draw())
		} else {
			fmt.Printf("Challenge successful.\n")
			t.Subject.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", t.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
			return
		}
	}

	t.Subject.Coins += 3

	Account(t.Subject)
}
