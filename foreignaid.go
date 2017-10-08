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

// func NewForeignAid(sub *Player) *Move {
// 	return NewMove(
// 		ForeignAid,
// 		sub,
// 		fmt.Sprintf("%s takes foreign aid.", sub.Name),
// 		0,
// 		nil,
// 		[]CardType{Duke},
// 		ForeignAidFunc(sub),
// 	)
// }
//
// func ForeignAidFunc(sub *Player) func() {
// 	return func() {
// 		sub.Coins += 2
// 		Account(sub)
// 	}
// }
