package coup

type Player struct {
	Name        string
	Chooser     Chooser
	Coins       int
	Hand        *Cards
	Placeholder bool
}

func NewPlayer(name string, chooser Chooser, placeholer bool) *Player {
	return &Player{
		Name:        name,
		Chooser:     chooser,
		Coins:       0,
		Hand:        NewCards(0, 0, 0, 0, 0),
		Placeholder: placeholer,
	}
}

func (p *Player) Valid(game *Game) []*Move {
	moves := []*Move{}

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

func (p *Player) Draw(deck *Cards) {
	card := deck.Peek()
	deck.Remove(card)
	p.Hand.Add(card)
}

func (p *Player) Discard(gm *Game, amt int) []CardEnum {
	return p.Chooser.ChooseDiscard(p.Hand, amt)
}

func (p *Player) Move(game *Game) *Move {
	return p.Chooser.ChooseMove(game, p.Valid(game))
}

func (p *Player) Block(game *Game, mv *Move) *Block {

	claims := []*Claim{}

	switch mv.Case {
	case ForeignAid:
		claims = append(claims, NewClaim(p, Duke))
	case Assassinate:
		claims = append(claims, NewClaim(p, Contessa))
	case Steal:
		claims = append(claims, NewClaim(p, Ambassador))
		claims = append(claims, NewClaim(p, Captain))
	}

	if claim := p.Chooser.ChooseBlock(claims); claim != nil {
		return NewBlock(p, claim)
	}

	return nil
}

func (p *Player) ChallengeMove(game *Game, claim *Claim, object *Player) *Challenge {
	if p.Chooser.ChooseChallengeMove(game, p, claim, object) {
		return NewChallenge(p)
	}
	return nil
}

func (p *Player) ChallengeBlock(game *Game, claim *Claim, object *Player) *Challenge {
	if p.Chooser.ChooseChallengeBlock(game, p, claim, object) {
		return NewChallenge(p)
	}
	return nil
}

func (p *Player) Observe(gm *Game, mv *Move, blk *Block, second bool) {
	p.Chooser.Update(p, gm, mv, blk, second)
}
