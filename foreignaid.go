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

func (f *ForeignAid) Modify(state *State) {
	f.Subject.Coins += 2
}

func (f *ForeignAid) Dispute() {}

func (f *ForeignAid) Impede() {

}

func (f *ForeignAid) DisputeImpedement() {

}

func (f *ForeignAid) Describe() {
	fmt.Printf("%s takes foreign aid.\n", f.Subject.Name)

	if f.Block != nil {

	}

	Account(f.Subject)
}
