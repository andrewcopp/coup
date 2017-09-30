package coup

import "fmt"

type ForeignAid struct {
	Subject *Player
	Block   *Block
}

func NewForeignAid(sub *Player) *Move {
	foreignAid := ForeignAid{
		Subject: sub,
	}

	return NewMove(
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
			block.Subject = other
			f.Block = block
			break
		}
	}

	if f.Block != nil {
		fmt.Printf("%s blocks foreign aid.\n", f.Block.Subject.Name)
	}

	f.Resolve(state)
}
