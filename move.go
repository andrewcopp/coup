package coup

import "fmt"

type Move struct {
	Announce func()
	Pay      func()
	Claim    *Claim
	Resolve  func(state *State)
}

func NewMove(announce func(), pay func(), claim *Claim, resolve func(state *State)) *Move {
	return &Move{
		Announce: announce,
		Pay:      pay,
		Claim:    claim,
		Resolve:  resolve,
	}
}

func (m *Move) Successful() bool {
	if m.Claim == nil {
		return true
	}

	return !m.Claim.Challenge.Successful
}

func Account(p *Player) {
	var format string
	if p.Coins != 1 {
		format = "%s has %d coins.\n"
	} else {
		format = "%s has %d coin.\n"
	}
	fmt.Printf(format, p.Name, p.Coins)
}
