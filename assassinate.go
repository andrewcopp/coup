package coup

type Assassinate struct {
	Subject *Player
	Object  *Player
}

func NewAssassinate(sub *Player, obj *Player) *Assassinate {
	return &Assassinate{
		Subject: sub,
		Object:  obj,
	}
}

func (a *Assassinate) Pay() {
	a.Subject.Coins -= 3
}

func (a *Assassinate) Claim() *Claim {
	return NewClaim(a.Subject, Assassin)
}

func (a *Assassinate) Counter() *func(game *Game) *Block {
	blockFunc := func(game *Game) *Block {
		claim := NewClaim(a.Object, Contessa)
		return a.Object.Block(game, claim)
	}
	return &blockFunc
}

func (a *Assassinate) Resolve() func(game *Game) {
	return func(game *Game) {
		if a.Object.Alive() {
			game.Board.Deck.Add(a.Object.Discard(a.Object.Chooser.ChooseDiscard()))
		}
	}
}
