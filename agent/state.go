package agent

type State struct {
	One         *Player
	Two         *Player
	Three       *Player
	Four        *Player
	Five        *Player
	Dukes       int
	Ambassadors int
	Assassins   int
	Captains    int
	Contessas   int
	Subject     *Player
	Action      Action
	Object      *Player
	Challenge   *Challenge
	Block       *Block
}

type Player struct {
	Coins int
	One   Card
	Two   Card
}

type Card int

const (
	None       Card = iota
	Duke            = iota
	Ambassador      = iota
	Assassin        = iota
	Captain         = iota
	Contessa        = iota
)

type Action int

const (
	Income      Action = iota
	ForeignAid         = iota
	Coup               = iota
	Tax                = iota
	Assassinate        = iota
	Exchange           = iota
	Steal              = iota
)

type Challenge struct {
	Subject    *Player
	Successful bool
}

type Block struct {
	Subject   *Player
	Card      Card
	Challenge *Challenge
}
