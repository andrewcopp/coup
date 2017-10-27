package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andrewcopp/coup"
)

func init() {

}

func main() {

	for i := 0; i < 5; i++ {
		wins := 0
		losses := 0
		var chooser coup.Chooser
		chooser = coup.NewAgent(i, 0.8)
		one := coup.NewPlayer("Player One", chooser, false)
		for j := 0; j < 1000; j++ {
			rand.Seed(time.Now().UnixNano())
			if i > 0 {
				chooser = coup.NewAgent(rand.Intn(i), 0.0)
			} else {
				chooser = coup.NewRandom()
			}
			two := coup.NewPlayer("Player Two", chooser, false)
			chooser = coup.NewRandom()
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
		fmt.Println()
		fmt.Println(float64(wins)/float64((wins+losses))*100.0, "%")
		fmt.Println()
	}

}
