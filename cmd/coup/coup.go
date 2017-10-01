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
	one := coup.NewPlayer("Player One", &decider, 2)
	two := coup.NewPlayer("Player Two", &decider, 2)
	three := coup.NewPlayer("Player Three", &decider, 2)
	four := coup.NewPlayer("Player Four", &decider, 2)
	five := coup.NewPlayer("Player Five", &decider, 2)

	board.Setup([]*coup.Player{one, two, three, four, five})
}

func main() {
	winner := board.Play()
	fmt.Println()
	fmt.Println(winner.Name, "wins!")
	fmt.Println()
}
