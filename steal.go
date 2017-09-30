package coup

import "fmt"

type Steal struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewSteal(sub *Player, obj *Player) *Steal {
	return &Steal{
		Subject: sub,
		Object:  obj,
	}
}

func (s *Steal) Modify(state *State) {

	fmt.Printf("%s steals from %s.\n", s.Subject.Name, s.Object.Name)

	claim := NewClaim(s.Subject, Captain, s.Object)
	for _, other := range state.Alive()[1:] {
		if claim.Challenger == nil {
			other.Dispute(claim)
		}
	}

	if claim.Challenger != nil {
		fmt.Printf("%s challenges.\n", claim.Challenger.Name)
		if card := s.Subject.Produce(claim.Declared); card != nil {
			fmt.Printf("Challenge unsuccessful.\n")
			claim.Challenger.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", claim.Challenger.Name, state.Revealed[len(state.Revealed)-1].Name())
			state.Deck.Add(card)
			s.Subject.Hand = append(s.Subject.Hand, state.Deck.Draw())
		} else {
			fmt.Printf("Challenge successful.\n")
			s.Subject.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", s.Subject.Name, state.Revealed[len(state.Revealed)-1].Name())
			return
		}
	}

	if s.Object.Coins == 0 {

	} else if s.Object.Coins == 1 {
		s.Object.Coins--
		s.Subject.Coins++
	} else {
		s.Object.Coins -= 2
		s.Object.Coins += 2
	}

	Account(s.Subject)
	Account(s.Object)

}
