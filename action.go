package coup

import "fmt"

type Action struct {
	Move  *Move
	Block *Block
}

func NewAction(mv *Move) *Action {
	return &Action{
		Move:  mv,
		Block: nil,
	}
}

func (a *Action) Apply(state *State) {
	fmt.Println(a.Move.Announcement)
	if a.Move.Cost != 0 {
		state.Players[0].Coins -= a.Move.Cost
		Account(state.Players[0])
	}

	if a.Move.Claim != nil {
		a.Move.Claim.Scrutinize(state)
	}

	if !a.Move.Successful() {
		return
	}

	a.Block = a.Move.Block(state)
	if a.Block != nil {
		fmt.Printf("%s blocks with a %s.\n", a.Block.Claim.Subject.Name, a.Block.Claim.Name())
		if a.Block.Successful() {
			return
		}
	}

	a.Move.Resolve()
}
