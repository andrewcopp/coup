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
		if block := s.Object.Block(game, NewClaim(Ambassador, nil)); block != nil {
			return block
		} else if block := s.Object.Block(game, NewClaim(Captain, nil)); block != nil {
			return block
		}
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
