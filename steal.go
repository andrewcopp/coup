package coup

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
//
// func StealFunc(sub *Player, obj *Player) func() {
// 	return func() {
// 		amt := 2
// 		if obj.Coins < 2 {
// 			amt = obj.Coins
// 		}
//
// 		obj.Coins -= amt
// 		sub.Coins += amt
//
// 		Account(sub)
// 		Account(obj)
// 	}
// }
