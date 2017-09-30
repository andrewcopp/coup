package coup

type Move struct {
	Action  *Action
	Counter *Counter
}

func NewMove(act *Action) *Move {
	return &Move{
		Action:  act,
		Counter: nil,
	}
}
