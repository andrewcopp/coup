package coup

import "fmt"

type ForeignAid struct {
	Subject *Player
	Block   *Block
}

func NewForeignAid(sub *Player) *Action {
	foreignAid := ForeignAid{
		Subject: sub,
	}

	return NewAction(
		foreignAid.Announce,
		foreignAid.Pay,
		foreignAid.Claim,
		foreignAid.Resolve,
	)
}

func (f *ForeignAid) Announce() {
	fmt.Printf("%s takes foreign aid.\n", f.Subject.Name)
}

func (f *ForeignAid) Pay() {}

func (f *ForeignAid) Claim(state *State) bool {
	return true
}

func (f *ForeignAid) Resolve(state *State) {
	f.Subject.Coins += 2
	Account(f.Subject)
}

func (f *ForeignAid) Modify(state *State) {

	others := state.Alive()[1:]
	for _, other := range others {
		if block := (*other.Brain).BlockForeignAid(state, f.Subject); block != nil {
			block.Blocker = other
			f.Block = block
			break
		}
	}

	if f.Block != nil {
		fmt.Printf("%s blocks foreign aid.\n", f.Block.Blocker.Name)
	}

	f.Resolve(state)
}
