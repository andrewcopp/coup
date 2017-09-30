package coup

import "fmt"

type Action struct {
	Move    *Move
	Counter *Counter
}

func NewAction(mv *Move) *Action {
	return &Action{
		Move:    mv,
		Counter: nil,
	}
}

func (a *Action) Apply(state *State) {
	a.Move.Announce()
	a.Move.Pay()

	if a.Move.Claim != nil {
		Make(a.Move.Claim, state)
	}

	if !a.Move.Successful() {
		return
	}

	a.Move.Resolve(state)

}

func Make(claim *Claim, state *State) {

	for _, other := range state.Alive()[1:] {
		if claim.Challenge.Subject == nil {
			other.Dispute(claim)
		}
	}

	if challenger := claim.Challenge.Subject; challenger != nil {
		fmt.Printf("%s challenges.\n", challenger.Name)
		if card := claim.Subject.Produce(claim.Declared); card != nil {
			claim.Challenge.Successful = false
			fmt.Printf("Challenge unsuccessful.\n")
			challenger.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
			state.Deck.Add(card)
			claim.Subject.Hand = append(claim.Subject.Hand, state.Deck.Draw())
		} else {
			fmt.Printf("Challenge successful.\n")
			claim.Subject.Reveal(state)
			claim.Challenge.Successful = true
			fmt.Printf("%s reveals a %s.\n", claim.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
		}
	}

}
