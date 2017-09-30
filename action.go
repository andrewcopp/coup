package coup

import "fmt"

type Action struct {
	Announce func()
	Pay      func()
	Claim    func(state *State) bool
	Resolve  func(state *State)
}

func (a *Action) Modify(state *State) {
	a.Announce()
	a.Pay()

	if !a.Claim(state) {
		return
	}

	a.Resolve(state)
}

func NewAction(announce func(), pay func(), claim func(state *State) bool, resolve func(state *State)) *Action {
	return &Action{
		Announce: announce,
		Pay:      pay,
		Claim:    claim,

		Resolve: resolve,
	}
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

type Challenge struct {
	Subject    *Player
	Successful bool
}

func NewChallenge(sub *Player) *Challenge {
	return &Challenge{
		Subject: sub,
	}
}

type Block struct {
	Subject   *Player
	Declared  Type
	Challenge *Challenge
}

func NewBlock() *Block {
	return &Block{}
}
