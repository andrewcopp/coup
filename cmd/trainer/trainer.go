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

	// if err := exec.Command("python3", "/Users/andrewcopp/Developer/Coup/initialize.py", "./cmd/trainer/models/model_1.cptk").Run(); err != nil {
	// 	fmt.Println(err)
	// }

	for i := 0; i < 1; i++ {
		wins := 0
		losses := 0
		var chooser coup.Chooser
		chooser = coup.NewAgent(i, 0.8)
		one := coup.NewPlayer("Player One", chooser, false)
		for j := 0; j < 1; j++ {

			// if err := exec.Command("python3", "/Users/andrewcopp/Developer/Coup/train.py", "./cmd/trainer/models/model_1.cptk", "./cmd/trainer/models/model_1.cptk", "4.0,2.0", "10.0").Run(); err != nil {
			// 	fmt.Println(err)
			// }
			//
			// if err := exec.Command("python3", "/Users/andrewcopp/Developer/Coup/train.py", "./cmd/trainer/models/model_1.cptk", "./cmd/trainer/models/model_1.cptk", "1.0,3.0", "5.0").Run(); err != nil {
			// 	fmt.Println(err)
			// }
			//
			// if err := exec.Command("python3", "/Users/andrewcopp/Developer/Coup/train.py", "./cmd/trainer/models/model_1.cptk", "./cmd/trainer/models/model_1.cptk", "2.0,3.0", "7.0").Run(); err != nil {
			// 	fmt.Println(err)
			// }

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
		fmt.Println(float64(wins)/float64((wins+losses))*100.0, "%")
	}

	// bits := binary.LittleEndian.Uint64(bytes)
	// float := math.Float64frombits(bits)
	// fmt.Println(float)

}
