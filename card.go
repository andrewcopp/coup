package coup

type Type int

const (
	Duke       Type = iota
	Ambassador      = iota
	Assassin        = iota
	Captain         = iota
	Contessa        = iota
)

type Card struct {
	Type Type
}

func NewCard(t Type) *Card {
	return &Card{
		Type: t,
	}
}

func (c *Card) Copy() *Card {
	return NewCard(c.Type)
}

func (c *Card) Name() string {
	switch c.Type {
	case Duke:
		return "Duke"
	case Ambassador:
		return "Ambassador"
	case Assassin:
		return "Assassin"
	case Captain:
		return "Captain"
	case Contessa:
		return "Contessa"
	}

	panic("Invalid Card.")
}
