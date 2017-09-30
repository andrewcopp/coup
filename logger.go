package coup

import "fmt"

func Account(p *Player) {
	var format string
	if p.Coins != 1 {
		format = "%s has %d coins.\n"
	} else {
		format = "%s has %d coin.\n"
	}
	fmt.Printf(format, p.Name, p.Coins)
}
