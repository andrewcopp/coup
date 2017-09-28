package coup

import "fmt"

type Action interface {
	Modify(state *State)
}

// Income
type Collect struct {
	Actor *Player
}

func NewCollect(actor *Player) *Collect {
	return &Collect{
		Actor: actor,
	}
}

func (c *Collect) Modify(state *State) {
	c.Actor.Coins++
	fmt.Println(c.Actor.Name, "takes income.")
}

// Foreign Aid
type Beg struct {
	Block *Block
}

func (b *Beg) Modify(state *State) {

}

// Coup
type Overthrow struct {
	Actor  *Player
	Target *Player
}

func NewOverthrow(actor *Player, target *Player) *Overthrow {
	return &Overthrow{
		Actor:  actor,
		Target: target,
	}
}

func (o *Overthrow) Modify(state *State) {
	o.Actor.Coins -= 7
	card := o.Target.Hidden[0]
	o.Target.Hidden = o.Target.Hidden[1:]
	o.Target.Revealed = append(o.Target.Revealed, card)
	fmt.Println(o.Target.Name, "reveals", card, ".")
}

type Tax struct {
	Executor  *Player
	Challenge *Challenge
}

type Assassinate struct {
	Executer  *Player
	Target    *Player
	Challenge *Challenge
	Block     *Block
}

type Exchange struct {
	Executer  *Player
	Challenge *Challenge
}

type Steal struct {
	Executer  *Player
	Target    *Player
	Challenge *Challenge
	Block     *Block
}

type Challenge struct {
	Challenger *Player
	Revealed   Card
}

type Block struct {
	Declared  Card
	Challenge *Challenge
}
