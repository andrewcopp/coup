package main

import (
	"fmt"

	"github.com/andrewcopp/coup"
)

var board *coup.Board

func init() {
	board = coup.NewBoard()

	var decider coup.Decider
	decider = coup.NewRandom()
	one := coup.NewPlayer("Player One", &decider, 2, []*coup.Card{}, []*coup.Card{})
	two := coup.NewPlayer("Player Two", &decider, 2, []*coup.Card{}, []*coup.Card{})
	three := coup.NewPlayer("Player Three", &decider, 2, []*coup.Card{}, []*coup.Card{})
	four := coup.NewPlayer("Player Four", &decider, 2, []*coup.Card{}, []*coup.Card{})

	board.Setup([]*coup.Player{one, two, three, four})
}

func main() {
	winner := board.Play()
	fmt.Println(winner.Name, "wins!")
}
