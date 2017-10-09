package coup

type Player struct {
	Name    string
	Chooser Chooser
	Coins   int
	Hand    *Hand
}

func NewPlayer(name string, chooser Chooser, coins int) *Player {
	return &Player{
		Name:    name,
		Chooser: chooser,
		Coins:   coins,
		Hand:    NewHand([]*Card{}),
	}
}

func (p *Player) Valid(game *Game) []Move {
	moves := []Move{}

	if game.Players[0].Coins < 10 {
		moves = append(moves, NewIncome(p))
		moves = append(moves, NewForeignAid(p))
		moves = append(moves, NewTax(p))
		moves = append(moves, NewExchange(p))
	}

	for _, other := range p.Opponents(game) {
		if game.Players[0].Coins >= 7 {
			moves = append(moves, NewCoup(p, other))
		}

		if game.Players[0].Coins < 10 {
			moves = append(moves, NewSteal(p, other))
		}

		if game.Players[0].Coins >= 3 {
			moves = append(moves, NewAssassinate(p, other))
		}
	}

	return moves
}

func (p *Player) Opponents(game *Game) []*Player {
	var index int
	for i, player := range game.Players {
		if p == player {
			index = i
		}
	}
	opponents := []*Player{}
	for _, player := range append(game.Players[index+1:], game.Players[:index]...) {
		if player.Alive() {
			opponents = append(opponents, player)
		}
	}
	return opponents
}

func (p *Player) Alive() bool {
	return p.Hand.Size() > 0
}

func (p *Player) Draw(deck *Deck) {
	p.Hand.Add(deck.Draw())
}

func (p *Player) Move(game *Game) Move {
	return p.Chooser.ChooseMove(p.Valid(game))
}

func (p *Player) Block(game *Game, claim *Claim) *Block {
	return nil
}

func (p *Player) Challenge(game *Game, claim *Claim) *Challenge {
	return nil
}
