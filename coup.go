package coup

type Coup struct {
	Subject *Player
	Object  *Player
}

func NewCoup(sub *Player, obj *Player) *Coup {
	return &Coup{
		Subject: sub,
		Object:  obj,
	}
}

func (c *Coup) Pay() {
	c.Subject.Coins -= 7
}

func (c *Coup) Claim() *Claim {
	return nil
}

func (c *Coup) Counter() *func(game *Game) *Block {
	return nil
}

func (c *Coup) Resolve() func(game *Game) {
	return func(game *Game) {
		game.Board.Deck.Add(c.Object.Discard(c.Object.Chooser.ChooseDiscard()))
	}
}
