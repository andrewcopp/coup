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

	if s.Object.Coins == 0 {

	} else if s.Object.Coins == 1 {
		s.Object.Coins--
		s.Subject.Coins++
	} else {
		s.Object.Coins -= 2
		s.Object.Coins += 2
	}

}

func (s *Steal) Dispute() {

}

func (s *Steal) Impede() {

}

func (s *Steal) DisputeImpedement() {

}

func (s *Steal) Describe() {
	fmt.Printf("%s steals from %s.\n", s.Subject.Name, s.Object.Name)
	Account(s.Subject)
	Account(s.Object)
}
