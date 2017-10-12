package main

import (
	"fmt"

	"github.com/andrewcopp/coup"
)

var game *coup.Game

func init() {
	var chooser coup.Chooser
	chooser = coup.NewAgent()
	one := coup.NewPlayer("Player One", chooser, 2)
	chooser = coup.NewRandom()
	two := coup.NewPlayer("Player Two", chooser, 2)
	// three := coup.NewPlayer("Player Three", chooser, 2)
	// four := coup.NewPlayer("Player Four", chooser, 2)
	// five := coup.NewPlayer("Player Five", chooser, 2)

	// game = coup.NewGame([]*coup.Player{one, two, three, four, five})
	game = coup.NewGame([]*coup.Player{one, two})
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
