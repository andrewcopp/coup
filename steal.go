package coup

import "fmt"

type Steal struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewSteal(sub *Player, obj *Player) *Move {
	steal := Steal{
		Subject: sub,
		Object:  obj,
	}

	return NewMove(
		steal.Announce,
		steal.Pay,
		steal.Claim,
		steal.Resolve,
	)
}

func (s *Steal) Announce() {
	fmt.Printf("%s steals from %s.\n", s.Subject.Name, s.Object.Name)
}

func (s *Steal) Pay() {}

func (s *Steal) Claim(state *State) bool {
	claim := NewClaim(s.Subject, Captain, s.Object)
	s.Subject.Make(claim, state)
	if claim.Revealed == nil {
		return true
	}

	if *claim.Revealed == claim.Declared {
		return true
	}

	return false
}

func (s *Steal) Resolve(state *State) {
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

func (s *Steal) Modify(state *State) {

}
