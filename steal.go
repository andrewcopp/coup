package coup

type Steal struct {
	Subject *Player
	Object  *Player
}

func NewSteal(sub *Player, obj *Player) *Steal {
	return &Steal{
		Subject: sub,
		Object:  obj,
	}
}

func (s *Steal) Pay() {}

func (s *Steal) Claim() *Claim {
	return NewClaim(Captain, s.Object)
}

func (s *Steal) Counter() *func(game *Game) *Block {
	blockFunc := func(game *Game) *Block {
		return nil
	}
	return &blockFunc
}

func (s *Steal) Resolve() {
	amt := 2
	if s.Object.Coins < 2 {
		amt = s.Object.Coins
	}

	s.Object.Coins -= amt
	s.Subject.Coins += amt
}

// func NewSteal(sub *Player, obj *Player) *Move {
// 	return NewMove(
// 		Steal,
// 		sub,
// 		fmt.Sprintf("%s steals from %s.", sub.Name, obj.Name),
// 		0,
// 		NewClaim(sub, Captain, obj),
// 		[]CardType{Ambassador, Captain},
// 		StealFunc(sub, obj),
// 	)
// }
