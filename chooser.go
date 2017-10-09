package coup

type Chooser interface {
	ChooseMove(moves []Move) Move
}
