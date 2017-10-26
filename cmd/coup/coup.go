package main

import (
	"fmt"

	"github.com/andrewcopp/coup"
)

var game *coup.Game

func init() {
	var chooser coup.Chooser
	chooser = coup.NewAgent()
	one := coup.NewPlayer("Player One", chooser, false)
	chooser = coup.NewRandom()
	two := coup.NewPlayer("Player Two", chooser, false)
	three := coup.NewPlayer("Player Three", chooser, false)
	four := coup.NewPlayer("Player Four", chooser, false)
	five := coup.NewPlayer("Player Five", chooser, false)

	game = coup.NewGame([]*coup.Player{one, two, three, four, five})
	// game = coup.NewGame([]*coup.Player{one, two})
}

func main() {
	game.Setup()
	winner := game.Play()
	if game.Logs {
		fmt.Println()
		fmt.Println(winner.Name, "wins!")
		fmt.Println()
	}
}
