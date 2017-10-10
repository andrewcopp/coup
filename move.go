package coup

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
	Case    MoveEnum
	Subject *Player
	Object  *Player
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

	switch m.Case {
	case Coup:
		m.Subject.Coins -= 7
	case Assassinate:
		m.Subject.Coins -= 3
	}

	if m.Exposed(gm) {
		return
	}

	if m.Blocked(gm) {
		return
	}

	switch m.Case {
	case Income:
		m.Subject.Coins += 1
	case ForeignAid:
		m.Subject.Coins += 2
	case Coup:
		m.Object.Discard(gm, 1)
	case Tax:
		m.Subject.Coins += 3
	case Assassinate:
		m.Object.Discard(gm, 1)
	case Exchange:
		// Draw 2
		m.Object.Discard(gm, 2)
	case Steal:
		amt := 2

		if m.Object.Coins < 2 {
			amt = m.Object.Coins
		}

		m.Subject.Coins += amt
		m.Object.Coins -= amt
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

	exposed := false
	var challenge *Challenge
	for _, other := range m.Subject.Opponents(gm) {
		if challenge = other.Challenge(gm, claim); challenge != nil {
			if claim.Verify() {
				challenge.Subject.Discard(gm, 1)
				//
			} else {
				m.Subject.Discard(gm, 1)
				exposed = true
			}
			break
		}
	}

	switch m.Case {
	case Tax, Assassinate, Exchange, Steal:
		// Update
	}

	return exposed
}

func (m *Move) Blocked(gm *Game) bool {
	var block *Block

	switch m.Case {
	case ForeignAid:
		for _, other := range m.Subject.Opponents(gm) {
			block = other.Block(gm, nil)
			if block != nil {
				break
			}
		}
	case Assassinate:
		block = m.Subject.Block(gm, nil)
	case Steal:
		block = m.Subject.Block(gm, nil)
	}

	if block != nil {

	}

	switch m.Case {
	case ForeignAid, Assassinate, Steal:
		// Update
	}

	return false
}
