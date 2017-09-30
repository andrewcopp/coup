package coup

import "fmt"

type Assassinate struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewAssassinate(sub *Player, obj *Player) *Action {

	assassinate := Assassinate{
		Subject:   sub,
		Object:    obj,
		Challenge: nil,
		Block:     nil,
	}

	return NewAction(
		assassinate.Announce,
		assassinate.Pay,
		assassinate.Claim,
		assassinate.Resolve,
	)
}

func (a *Assassinate) Announce() {
	fmt.Printf("%s assassinates %s.\n", a.Subject.Name, a.Object.Name)
}

func (a *Assassinate) Pay() {
	a.Subject.Coins -= 3
	Account(a.Subject)
}

func (a *Assassinate) Claim(state *State) bool {
	claim := NewClaim(a.Subject, Assassin, a.Object)
	a.Subject.Make(claim, state)
	if claim.Revealed == nil {
		return true
	}

	if *claim.Revealed == claim.Declared {
		return true
	}

	return false
}

func (a *Assassinate) Resolve(state *State) {
	if len(a.Object.Hand) != 0 {
		a.Object.Reveal(state)
		fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
	}
}

func (a *Assassinate) Modify(state *State) {

}
