package coup

type MoveEnum int

const (
	Income      MoveEnum = iota
	ForeignAid           = iota
	Coup                 = iota
	Tax                  = iota
	Assassinate          = iota
	Exchange             = iota
	Steal                = iota
)

type Move struct {
	MoveEnum MoveEnum
	Subject  int // *Player
	Object   int
}

func NewIncome(sub int) *Move {
	return &Move{
		MoveEnum: Income,
		Subject:  sub,
	}
}

func NewForeignAid(sub int) *Move {
	return &Move{
		MoveEnum: ForeignAid,
		Subject:  sub,
	}
}

func NewCoup(sub int, obj int) *Move {
	return &Move{
		MoveEnum: Coup,
		Subject:  sub,
		Object:   obj,
	}
}

func NewTax(sub int) *Move {
	return &Move{
		MoveEnum: Tax,
		Subject:  sub,
	}
}

func NewAssassinate(sub int, obj int) *Move {
	return &Move{
		MoveEnum: Assassinate,
		Subject:  sub,
		Object:   obj,
	}
}

func NewExchange(sub int) *Move {
	return &Move{
		MoveEnum: Exchange,
		Subject:  sub,
	}
}

func NewSteal(sub int, obj int) *Move {
	return &Move{
		MoveEnum: Steal,
		Subject:  sub,
		Object:   obj,
	}
}
