package coup

import "fmt"

type Assassinate struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewAssassinate(sub *Player, obj *Player) *Assassinate {
	return &Assassinate{
		Subject:   sub,
		Object:    obj,
		Challenge: nil,
		Block:     nil,
	}
}

func (a *Assassinate) Modify(state *State) {
	fmt.Printf("%s assassinates %s.\n", a.Subject.Name, a.Object.Name)
	a.Subject.Coins -= 3
	Account(a.Subject)

	claim := NewClaim(a.Subject, Assassin, a.Object)
	for _, other := range state.Alive()[1:] {
		if claim.Challenger == nil {
			other.Dispute(claim)
		}
	}

	if claim.Challenger != nil {
		fmt.Printf("%s challenges.\n", claim.Challenger.Name)
		if card := a.Subject.Produce(claim.Declared); card != nil {
			fmt.Printf("Challenge unsuccessful.\n")
			claim.Challenger.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", claim.Challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
			state.Deck.Add(card)
			a.Subject.Hand = append(a.Subject.Hand, state.Deck.Draw())
		} else {
			fmt.Printf("Challenge successful.\n")
			a.Subject.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", a.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
			return
		}
	}

	if len(a.Object.Hand) != 0 {
		a.Object.Reveal(state)
		fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
	}

}
