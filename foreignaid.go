package coup

import "fmt"

type ForeignAid struct {
	Subject *Player
	Block   *Block
}

func NewForeignAid(sub *Player) *ForeignAid {
	return &ForeignAid{
		Subject: sub,
	}
}

func (f *ForeignAid) Announce() {
	fmt.Printf("%s takes foreign aid.\n", f.Subject.Name)
}

func (f *ForeignAid) Modify(state *State) {
	f.Subject.Coins += 2

	others := state.Alive()[1:]
	for _, other := range others {
		if block := (*other.Brain).BlockForeignAid(state, f.Subject); block != nil {
			block.Blocker = other
			f.Block = block
			f.Subject.Coins -= 2
			break
		}
	}

	if f.Block != nil {
		fmt.Printf("%s blocks foreign aid.\n", f.Block.Blocker.Name)
	}

	Account(f.Subject)
}
