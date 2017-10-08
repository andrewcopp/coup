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

// func NewIncome(sub *Player) *Move {
// 	return NewMove(
// 		Income,
// 		sub,
// 		fmt.Sprintf("%s takes income.", sub.Name),
// 		0,
// 		nil,
// 		[]CardType{},
// 		IncomeFunc(sub),
// 	)
// }
//
// func IncomeFunc(sub *Player) func() {
// 	return func() {
// 		sub.Coins++
// 		Account(sub)
// 	}
// }
