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
	a.Subject.Make(claim, state)
	if claim.Revealed == nil {
		if len(a.Object.Hand) != 0 {
			a.Object.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
		}
		return
	}

	if *claim.Revealed == claim.Declared {
		if len(a.Object.Hand) != 0 {
			a.Object.Reveal(state)
			fmt.Printf("%s reveals a %s.\n", a.Object.Name, state.Revealed[len(state.Revealed)-1].Name())
		}
		return
	}

}
