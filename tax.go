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

func (t *Tax) Announce() {
	fmt.Printf("%s taxes.\n", t.Subject.Name)
}

func (t *Tax) Modify(state *State) {

	claim := NewClaim(t.Subject, Duke, nil)
	t.Subject.Make(claim, state)
	if claim.Revealed == nil {
		t.Subject.Coins += 3
		Account(t.Subject)
		return
	}

	if *claim.Revealed == claim.Declared {
		t.Subject.Coins += 3
		Account(t.Subject)
		return
	}

}
