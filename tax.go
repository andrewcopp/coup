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

func (t *Tax) Pay() {}

func (t *Tax) Claim(state *State) bool {
	claim := NewClaim(t.Subject, Duke, nil)
	t.Subject.Make(claim, state)
	if claim.Revealed == nil {
		return true
	}

	if *claim.Revealed == claim.Declared {
		return true
	}

	return false
}

func (t *Tax) Resolve(state *State) {
	t.Subject.Coins += 3
	Account(t.Subject)
}

func (t *Tax) Modify(state *State) {

	claim := NewClaim(t.Subject, Duke, nil)
	t.Subject.Make(claim, state)
	if claim.Revealed == nil {
		t.Resolve(state)
		return
	}

	if *claim.Revealed == claim.Declared {
		t.Resolve(state)
		return
	}

}
