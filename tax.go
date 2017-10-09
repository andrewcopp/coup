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
	return NewClaim(t.Subject, Duke)
}

func (t *Tax) Counter() *func(game *Game) *Block {
	return nil
}

func (t *Tax) Resolve() {
	t.Subject.Coins += 3
}
