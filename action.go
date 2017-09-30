package coup

type Action struct {
	Move    *Move
	Counter *Counter
}

func NewAction(mv *Move) *Action {
	return &Action{
		Move:    mv,
		Counter: nil,
	}
}
