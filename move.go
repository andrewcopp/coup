package coup

type Move interface {
	Pay()
	Claim() *Claim
	Counter() *func(game *Game) *Block
	Resolve()
}
