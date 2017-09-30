package coup

type Block struct {
	Subject   *Player
	Declared  Type
	Challenge *Challenge
}

func NewBlock(sub *Player, dec Type) *Block {
	return &Block{
		Subject:  sub,
		Declared: dec,
	}
}
