package coup

// type ForeignAid struct {
// 	Subject *Player
// }
//
// func NewForeignAid(sub *Player) *ForeignAid {
// 	return &ForeignAid{
// 		Subject: sub,
// 	}
// }
//
// func (f *ForeignAid) Pay() {}
//
// func (f *ForeignAid) Claim() *Claim {
// 	return nil
// }
//
// func (f *ForeignAid) Counter() *func(game *Game) *Block {
// 	blockFunc := func(game *Game) *Block {
// 		for _, opponent := range f.Subject.Opponents(game) {
// 			claim := NewClaim(f.Subject, Duke)
// 			if block := opponent.Block(game, claim); block != nil {
// 				return block
// 			}
// 		}
// 		return nil
// 	}
// 	return &blockFunc
// }
//
// func (f *ForeignAid) Resolve() func(game *Game) {
// 	return func(game *Game) {
// 		f.Subject.Coins += 2
// 	}
// }
