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
	return NewClaim(s.Subject, Captain)
}

func (s *Steal) Counter() *func(game *Game) *Block {
	blockFunc := func(game *Game) *Block {
		if block := s.Object.Block(game, NewClaim(s.Object, Ambassador)); block != nil {
			return block
		} else if block := s.Object.Block(game, NewClaim(s.Object, Captain)); block != nil {
			return block
		}
		return nil
	}
	return &blockFunc
}

func (s *Steal) Resolve() func(game *Game) {
	return func(game *Game) {
		amt := 2
		if s.Object.Coins < 2 {
			amt = s.Object.Coins
		}

		s.Object.Coins -= amt
		s.Subject.Coins += amt
	}
}
