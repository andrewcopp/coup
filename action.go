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

	a.Block = a.Move.Block()

	if claim := a.Move.Claim; claim != nil {
		if a.Move.Claim.Object != nil {
			if block := a.Move.Claim.Object.Impede(a.Move.Counters); block != nil {
				a.Block = block
				fmt.Printf("%s blocks with a %d.\n", block.Claim.Subject.Name, block.Claim.Declared)
				a.Block.Claim.Scrutinize(state)
				if a.Block.Successful() {
					return
				}
			}
		} else {
			for _, player := range a.Move.Claim.Subject.Opponents(state) {
				if block := player.Impede(a.Move.Counters); block != nil {
					a.Block = block
					fmt.Printf("%s blocks with a %d.\n", player.Name, block.Claim.Declared)
					a.Block.Claim.Scrutinize(state)
					if a.Block.Successful() {
						return
					}
				}
			}
		}
	}

	a.Move.Resolve(state)
}
