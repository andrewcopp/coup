package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/andrewcopp/coup"
)

func init() {

}

func main() {

	for i := 0; i < 10; i++ {

		if i != 0 {
			infile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", i-1)
			outfile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", i)
			if err := exec.Command("python3", "/home/ubuntu/reinforcement/transfer.py", infile, outfile).Run(); err != nil {
				fmt.Println(err)
			}
		} else {
			outfile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", i)
			if err := exec.Command("python3", "/home/ubuntu/reinforcement/initialize.py", outfile).Run(); err != nil {
				fmt.Println(err)
			}
		}

		wins := 0
		losses := 0

		bump := make(chan struct{}, 1000)

		for j := 0; j < 1000; j++ {
			go func() {
				var chooser coup.Chooser
				chooser = coup.NewAgent(i, 1.0/float64(i))
				one := coup.NewPlayer("Player One", chooser, false)

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
					one.Chooser.Record(true)
					wins++
				} else {
					one.Chooser.Record(false)
					losses++
				}
				bump <- struct{}{}
			}()
		}

		for j := 0; j < 1000; j++ {
			<-bump
			fmt.Printf("Agent %d completed game %d.\n", i, j)
		}

		fmt.Println(float64(wins)/float64((wins+losses))*100.0, "%")
	}

}
