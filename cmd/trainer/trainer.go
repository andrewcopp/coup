package main

import (
	"fmt"

	"github.com/andrewcopp/coup"
)

func init() {

}

func main() {

	var chooser coup.Chooser
	chooser = coup.NewAgent()
	one := coup.NewPlayer("Player One", chooser, false)

	wins := 0
	losses := 0
	for i := 0; i < 1; i++ {
		for j := 0; j < 2000; j++ {
			chooser = coup.NewRandom()
			two := coup.NewPlayer("Player Two", chooser, false)
			three := coup.NewPlayer("Player Three", chooser, true)
			four := coup.NewPlayer("Player Four", chooser, true)
			five := coup.NewPlayer("Player Five", chooser, true)

			game := coup.NewGame([]*coup.Player{one, two, three, four, five})

			game.Setup()
			winner := game.Play()
			if winner == one {
				wins++
			} else {
				losses++
			}
		}
	}

	fmt.Println()
	fmt.Println(float64(wins)/float64((wins+losses))*100.0, "%")
	fmt.Println()

}
