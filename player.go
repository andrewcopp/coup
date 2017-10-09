package coup

type Player struct {
	Name string
	// Brain *Decider
	Coins int
	Hand  *Hand
}

func (p *Player) Copy() *Player {
	return &Player{
		Name:  p.Name,
		Coins: p.Coins,
		Hand:  p.Hand.Copy(),
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

//
// func NewPlayer(name string, brain *Decider, coins int) *Player {
// 	return &Player{
// 		Name: name,
// 		// Brain: brain,
// 		Coins: coins,
// 		Hand:  []*Card{},
// 	}
// }

func (p *Player) Choose(moves []Move) Move {
	return moves[0]
}

func (p *Player) Opponents(game *Game) []*Player {
	opponents := []*Player{}
	for _, player := range game.Players {
		if player.Hand.Size() != 0 && p != player {
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
	return p.Choose(p.Valid(game))
}

func (p *Player) Block(game *Game, move Move) *Block {
	return nil
}

func (p *Player) Challenge(game *Game, claim *Claim) *Challenge {
	return nil
}
