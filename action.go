package coup

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
	a.Move.Announce()
	if a.Move.Cost != 0 {
		state.Players[0].Coins -= a.Move.Cost
		Account(state.Players[0])
	}

	if a.Move.Claim != nil {
		a.Move.Scrutinize(state)
	}

	if !a.Move.Successful() {
		return
	}

	a.Move.Resolve(state)
}
