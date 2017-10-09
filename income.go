package coup

type Income struct {
	Subject *Player
}

func NewIncome(sub *Player) *Income {
	return &Income{
		Subject: sub,
	}
}

func (i *Income) Pay() {}

func (i *Income) Claim() *Claim {
	return nil
}

func (i *Income) Counter() *func(game *Game) *Block {
	return nil
}

func (i *Income) Resolve() {
	i.Subject.Coins++
}
