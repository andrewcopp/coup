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
	// others := state.Alive()[1:]
	//
	// var challenger *Player
	// for _, other := range others {
	// 	challenge := (*other.Brain).ChallengeTax(state, other)
	// 	if challenge != nil {
	// 		challenger = other
	// 		break
	// 	}
	// }
	//
	// if challenger != nil {
	// 	card := t.Subject.Produce(Duke)
	// 	if card != nil {
	t.Subject.Coins += 3
	// 		challenger.Reveal()
	// 		state.Deck.Add(card)
	// 		t.Subject.Hidden = append(t.Subject.Hidden, state.Deck.Draw())
	// 		fmt.Println(challenger.Name, "unsuccessfully challenged.")
	// 		fmt.Println(challenger.Name, "reveals", challenger.Revealed[len(challenger.Revealed)-1].Name(), ".")
	// 		fmt.Println(t.Subject.Name, "taxes. Has", t.Subject.Coins, "coins.")
	// 	} else {
	// 		t.Subject.Reveal()
	// 		fmt.Println(challenger.Name, "successfully challenged.")
	// 		fmt.Println(t.Subject.Name, "reveals", t.Subject.Revealed[len(t.Subject.Revealed)-1].Name(), ".")
	// 	}
	// } else {
	// 	t.Subject.Coins += 3
	// 	fmt.Println(t.Subject.Name, "taxes. Has", t.Subject.Coins, "coins.")
	// }

}

func (t *Tax) Dispute(state *State) {
	_ = NewClaim(t.Subject, Duke, nil)
}

func (t *Tax) Impede(state *State) {}

func (t *Tax) Describe() {
	fmt.Printf("%s taxes.\n", t.Subject.Name)
	Account(t.Subject)
}
