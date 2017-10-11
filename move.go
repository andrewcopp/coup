package coup

import "fmt"

type MoveEnum int

const (
	Income      MoveEnum = iota
	ForeignAid           = iota
	Coup                 = iota
	Tax                  = iota
	Assassinate          = iota
	Exchange             = iota
	Steal                = iota
)

type Move struct {
	Case      MoveEnum
	Subject   *Player
	Object    *Player
	Challenge *Challenge
	Block     *Block
}

func NewIncome(sub *Player) *Move {
	return &Move{
		Case:    Income,
		Subject: sub,
	}
}

func NewForeignAid(sub *Player) *Move {
	return &Move{
		Case:    ForeignAid,
		Subject: sub,
	}
}

func NewCoup(sub *Player, obj *Player) *Move {
	return &Move{
		Case:    Coup,
		Subject: sub,
		Object:  obj,
	}
}

func NewTax(sub *Player) *Move {
	return &Move{
		Case:    Tax,
		Subject: sub,
	}
}

func NewAssassinate(sub *Player, obj *Player) *Move {
	return &Move{
		Case:    Assassinate,
		Subject: sub,
		Object:  obj,
	}
}

func NewExchange(sub *Player) *Move {
	return &Move{
		Case:    Exchange,
		Subject: sub,
	}
}

func NewSteal(sub *Player, obj *Player) *Move {
	return &Move{
		Case:    Steal,
		Subject: sub,
		Object:  obj,
	}
}

func (m *Move) Modify(gm *Game) {

	// TODO: Return this
	for _, player := range gm.Players {
		if player.Hand.Size() > 2 {
			for _, card := range player.Discard(gm, 2) {
				gm.Deck.Add(card)
			}
		}
	}

	switch m.Case {
	case Coup:
		m.Subject.Coins -= 7
	case Assassinate:
		m.Subject.Coins -= 3
	}

	if gm.Logs {
		switch m.Case {
		case Income:
			fmt.Printf("%s takes income.\n", m.Subject.Name)
		case ForeignAid:
			fmt.Printf("%s takes foreign aid.\n", m.Subject.Name)
		case Coup:
			fmt.Printf("%s coups %s.\n", m.Subject.Name, m.Object.Name)
		case Tax:
			fmt.Printf("%s taxes.\n", m.Subject.Name)
		case Assassinate:
			fmt.Printf("%s assassinates %s.\n", m.Subject.Name, m.Object.Name)
		case Exchange:
			fmt.Printf("%s exchanges.\n", m.Subject.Name)
		case Steal:
			fmt.Printf("%s steals from %s.\n", m.Subject.Name, m.Object.Name)
		}
	}

	exposed := m.Exposed(gm)

	if !exposed && m.Case == Exchange {
		// TODO: return this home
		if m.Case == Exchange {
			m.Subject.Draw(gm.Deck)
			m.Subject.Draw(gm.Deck)
		}
	}

	switch m.Case {
	case Tax, Assassinate, Exchange, Steal:
		var challenger *Player
		if m.Challenge != nil {
			challenger = m.Challenge.Subject
		}
		for _, player := range gm.Players {
			player.Observe(gm, m, challenger, nil, nil)
		}
	}

	if exposed {
		return
	}

	if m.Blocked(gm) {
		var blocker *Player
		var challenger *Player
		if m.Block != nil {
			blocker = m.Block.Subject
			if m.Block.Challenge != nil {
				challenger = m.Block.Challenge.Subject
			}
		}
		for _, player := range gm.Players {
			player.Observe(gm, nil, nil, blocker, challenger)
		}

		return
	}

	switch m.Case {
	case Income:
		m.Subject.Coins++
	case ForeignAid:
		m.Subject.Coins += 2
	case Coup:
		for _, card := range m.Object.Discard(gm, 1) {
			gm.Discard.Add(card)
		}
	case Tax:
		m.Subject.Coins += 3
	case Assassinate:
		if m.Object.Alive() {
			for _, card := range m.Object.Discard(gm, 1) {
				gm.Discard.Add(card)
			}
		}
	case Exchange:
		// TODO: return to here
	case Steal:
		amt := 2

		if m.Object.Coins < 2 {
			amt = m.Object.Coins
		}

		m.Subject.Coins += amt
		m.Object.Coins -= amt
	}

	if gm.Logs {
		switch m.Case {
		case Income:
			fmt.Printf("%s has %d coins.\n", m.Subject.Name, m.Subject.Coins)
		case ForeignAid:
			fmt.Printf("%s has %d coins.\n", m.Subject.Name, m.Subject.Coins)
		case Coup:
			fmt.Printf("%s discards.\n", m.Object.Name)
		case Tax:
			fmt.Printf("%s has %d coins.\n", m.Subject.Name, m.Subject.Coins)
		case Assassinate:
			fmt.Printf("%s potentially discards.\n", m.Object.Name)
		case Exchange:
			fmt.Printf("%s draws two cards.\n", m.Subject.Name)
		case Steal:
			fmt.Printf("%s has %d coins.\n", m.Subject.Name, m.Subject.Coins)
			fmt.Printf("%s has %d coins.\n", m.Object.Name, m.Object.Coins)
		}
	}

	var blocker *Player
	var challenger *Player
	if m.Block != nil {
		blocker = m.Block.Subject
		if m.Block.Challenge != nil {
			challenger = m.Block.Challenge.Subject
		}
	}
	for _, player := range gm.Players {
		player.Observe(gm, nil, nil, blocker, challenger)
	}

}

func (m *Move) Exposed(gm *Game) bool {
	var claim *Claim
	switch m.Case {
	case Tax:
		claim = NewClaim(m.Subject, Duke)
	case Assassinate:
		claim = NewClaim(m.Subject, Assassin)
	case Exchange:
		claim = NewClaim(m.Subject, Ambassador)
	case Steal:
		claim = NewClaim(m.Subject, Captain)
	}

	if claim == nil {
		return false
	}

	var exposed bool
	for _, other := range m.Subject.Opponents(gm) {
		if m.Challenge = other.Challenge(gm, claim); m.Challenge != nil {
			if gm.Logs {
				fmt.Printf("%s challenges.\n", m.Challenge.Subject.Name)
			}
			if claim.Verify() {
				if gm.Logs {
					fmt.Printf("Challenge unsuccessful.\n")
				}
				for _, card := range m.Challenge.Subject.Discard(gm, 1) {
					gm.Discard.Add(card)
				}
				m.Subject.Hand.Remove(claim.Declared)
				gm.Deck.Add(claim.Declared)
				card := gm.Deck.Peek()
				gm.Deck.Remove(card)
				m.Subject.Hand.Add(card)
			} else {
				if gm.Logs {
					fmt.Printf("Challenge successful.\n")
				}
				for _, card := range m.Subject.Discard(gm, 1) {
					gm.Discard.Add(card)
				}
				exposed = true
			}
			break
		}
	}

	return exposed
}

func (m *Move) Blocked(gm *Game) bool {
	switch m.Case {
	case ForeignAid:
		for _, other := range m.Subject.Opponents(gm) {
			m.Block = other.Block(gm, m)
			if m.Block != nil {
				break
			}
		}
	case Assassinate:
		if m.Object.Alive() {
			m.Block = m.Object.Block(gm, m)
		}
	case Steal:
		if m.Object.Alive() {
			m.Block = m.Object.Block(gm, m)
		}
	}

	var blocked bool
	if m.Block != nil {
		if gm.Logs {
			fmt.Printf("%s blocks.\n", m.Block.Subject.Name)
		}
		if !m.Block.Exposed(gm) {
			blocked = true
		}
	}

	return blocked
}
