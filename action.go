package coup

import "fmt"

type Action interface {
	Modify(state *State)
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
	Challenger *Player
	Revealed   Type
}

func NewChallenge(challenger *Player) *Challenge {
	return &Challenge{
		Challenger: challenger,
	}
}

type Block struct {
	Blocker   *Player
	Declared  Type
	Challenge *Challenge
}

func NewBlock() *Block {
	return &Block{}
}
