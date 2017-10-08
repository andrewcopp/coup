package coup

type Tax struct {
	Subject *Player
}

func NewTax(sub *Player) *Tax {
	return &Tax{
		Subject: sub,
	}
}

func (t *Tax) Pay() {}

func (t *Tax) Claim() *Claim {
	return NewClaim(Duke, nil)
}

func (t *Tax) Counter() *func(game *Game) *Block {
	return nil
}

func (t *Tax) Resolve() {
	t.Subject.Coins += 3
}

// func NewTax(sub *Player) *Move {
// 	return NewMove(
// 		Tax,
// 		sub,
// 		fmt.Sprintf("%s taxes.", sub.Name),
// 		0,
// 		NewClaim(sub, Duke, nil),
// 		[]CardType{},
// 		TaxFunc(sub),
// 	)
// }
