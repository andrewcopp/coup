package coup

type ForeignAid struct {
	Subject *Player
}

func NewForeignAid(sub *Player) *ForeignAid {
	return &ForeignAid{
		Subject: sub,
	}
}

func (f *ForeignAid) Pay() {}

func (f *ForeignAid) Claim() *Claim {
	return nil
}

func (f *ForeignAid) Counter() *func(game *Game) *Block {
	blockFunc := func(game *Game) *Block {
		return nil
	}
	return &blockFunc
}

func (f *ForeignAid) Resolve() {
	f.Subject.Coins += 2
}
