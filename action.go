package coup

import "fmt"

type Action interface {
	Modify(state *State)
}

type Income struct {
	Subject *Player
}

func NewIncome(sub *Player) *Income {
	return &Income{
		Subject: sub,
	}
}

func (i *Income) Modify(state *State) {
	i.Subject.Coins++
	fmt.Println(i.Subject.Name, "takes income.", i.Subject.Name, "has", i.Subject.Coins, "coins.")
}

type ForeignAid struct {
	Subject *Player
	Block   *Block
}

func NewForeignAid(sub *Player) *ForeignAid {
	return &ForeignAid{
		Subject: sub,
	}
}

func (f *ForeignAid) Modify(state *State) {
	fmt.Println(f.Subject.Name, "takes foreign aid.")
	f.Subject.Coins += 2
	fmt.Println(f.Subject.Name, "has", f.Subject.Coins, "coins.")
}

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

func (c *Coup) Modify(state *State) {
	c.Subject.Coins -= 7
	card := c.Object.Hidden[0]
	c.Object.Hidden = c.Object.Hidden[1:]
	c.Object.Revealed = append(c.Object.Revealed, card)
	fmt.Println(c.Subject.Name, "pays 7 coins to coup", c.Object.Name, ".", c.Subject.Name, "has", c.Subject.Coins, "coins.")
	fmt.Println(c.Object.Name, "reveals", card.Name(), ".")
}

type Tax struct {
	Subject   *Player
	Challenge *Challenge
}

func NewTax(sub *Player) *Tax {
	return &Tax{
		Subject:   sub,
		Challenge: nil,
	}
}

func (t *Tax) Modify(state *State) {
	others := state.Alive()[1:]

	var challenger *Player
	for _, other := range others {
		challenge := (*other.Brain).ChallengeTax(state, other)
		if challenge != nil {
			challenger = other
			break
		}
	}

	if challenger != nil {
		card := t.Subject.Produce(Duke)
		if card != nil {
			t.Subject.Coins += 3
			challenger.Reveal()
			state.Deck.Add(card)
			t.Subject.Hidden = append(t.Subject.Hidden, state.Deck.Draw())
			fmt.Println(challenger.Name, "unsuccessfully challenged.")
			fmt.Println(challenger.Name, "reveals", challenger.Revealed[len(challenger.Revealed)-1].Name(), ".")
			fmt.Println(t.Subject.Name, "taxes. Has", t.Subject.Coins, "coins.")
		} else {
			t.Subject.Reveal()
			fmt.Println(challenger.Name, "successfully challenged.")
			fmt.Println(t.Subject.Name, "reveals", t.Subject.Revealed[len(t.Subject.Revealed)-1].Name(), ".")
		}
	} else {
		t.Subject.Coins += 3
		fmt.Println(t.Subject.Name, "taxes. Has", t.Subject.Coins, "coins.")
	}

}

type Assassinate struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewAssassinate(sub *Player, obj *Player) *Assassinate {
	return &Assassinate{
		Subject:   sub,
		Object:    obj,
		Challenge: nil,
		Block:     nil,
	}
}

func (a *Assassinate) Modify(state *State) {
	a.Subject.Coins -= 3
	a.Object.Reveal()
	fmt.Println(a.Subject.Name, "assassinates", a.Object.Name, ".", a.Subject.Name, "has", a.Subject.Coins, "coins.")
	fmt.Println(a.Object.Name, "reveals", a.Object.Revealed[len(a.Object.Revealed)-1].Name(), ".")
}

type Exchange struct {
	Subject   *Player
	Challenge *Challenge
}

func NewExchange(sub *Player) *Exchange {
	return &Exchange{
		Subject: sub,
	}
}

func (e *Exchange) Modify(state *State) {
	e.Subject.Hidden = append(e.Subject.Hidden, state.Deck.Draw())
	e.Subject.Hidden = append(e.Subject.Hidden, state.Deck.Draw())

	state.Deck.Add(e.Subject.Discard())
	state.Deck.Add(e.Subject.Discard())

	fmt.Println(state.Players[0].Name, "exchanges with the deck.")
}

type Steal struct {
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewSteal(sub *Player, obj *Player) *Steal {
	return &Steal{
		Subject: sub,
		Object:  obj,
	}
}

func (s *Steal) Modify(state *State) {
	if s.Object.Coins == 0 {
		fmt.Println(s.Subject.Name, "steals 0 coins from", s.Object.Name, ".")
	} else if s.Object.Coins == 1 {
		s.Object.Coins--
		s.Subject.Coins++
		fmt.Println(s.Subject.Name, "steals 1 coins from", s.Object.Name, ".")
	} else {
		s.Object.Coins -= 2
		s.Object.Coins += 2
		fmt.Println(s.Subject.Name, "steals 2 coins from", s.Object.Name, ".")
	}

}

type Challenge struct {
	Challenger *Player
	Revealed   Type
}

func NewChallenge(challenger *Player) *Challenge {
	return &Challenge{
		Challenger: challenger,
	}
}

type Block struct {
	Declared  Type
	Challenge *Challenge
}
