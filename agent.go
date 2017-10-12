package coup

import "fmt"

type Agent struct {
	States []State
	Action Action
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Update(self *Player, gm *Game, mv *Move, blk *Block, second bool) {

	index := 0
	for i, player := range gm.Players {
		if self == player {
			index = i
		}
	}

	players := append(gm.Players[index:], gm.Players[:index]...)
	state := State{
		OneCoins:             float64(players[0].Coins) / 12.0,
		OneDukes:             float64(players[0].Hand.Dukes) / 3.0,
		OneAssassins:         float64(players[0].Hand.Assassins) / 3.0,
		OneAmbassadors:       float64(players[0].Hand.Ambassadors) / 3.0,
		OneCaptains:          float64(players[0].Hand.Captains) / 3.0,
		OneContessas:         float64(players[0].Hand.Contessas) / 3.0,
		TwoCoins:             float64(players[1].Coins) / 12.0,
		TwoCards:             float64(players[1].Hand.Size()) / 4.0,
		DiscardedDukes:       float64(gm.Discard.Dukes) / 3.0,
		DiscardedAssassins:   float64(gm.Discard.Assassins) / 3.0,
		DiscardedAmbassadors: float64(gm.Discard.Ambassadors) / 3.0,
		DiscardedCaptains:    float64(gm.Discard.Captains) / 3.0,
		DiscardedContessas:   float64(gm.Discard.Contessas) / 3.0,
		Bridge:               &Bridge{},
	}

	cap := len(gm.Players)

	one := index
	two := index + 1
	if two >= cap {
		two -= cap
	}

	if mv != nil {
		switch mv.Subject {
		case gm.Players[one]:
			state.Bridge.SubjectOne = true
		case gm.Players[two]:
			state.Bridge.SubjectTwo = true
		}

		switch mv.Case {
		case ForeignAid:
			state.Bridge.MoveForeignAid = true
		case Tax:
			state.Bridge.MoveTax = true
		case Assassinate:
			state.Bridge.MoveAssassinate = true
		case Exchange:
			state.Bridge.MoveExchange = true
		case Steal:
			state.Bridge.MoveSteal = true
		}

		if mv.Object != nil {
			switch mv.Object {
			case gm.Players[one]:
				state.Bridge.ObjectOne = true
			case gm.Players[two]:
				state.Bridge.ObjectTwo = true
			}
		}

		if mv.Challenge != nil {
			switch mv.Challenge.Subject {
			case gm.Players[one]:
				state.Bridge.MoveChallengeSubjectOne = true
			case gm.Players[two]:
				state.Bridge.MoveChallengeSubjectTwo = true
			}
		}
	}

	if blk != nil {
		switch blk.Subject {
		case gm.Players[one]:
			state.Bridge.BlockSubjectOne = true
		case gm.Players[two]:
			state.Bridge.BlockSubjectTwo = true
		}

		switch blk.Claim.Declared {
		case Ambassador:
			state.Bridge.BlockAmbassador = true
		case Captain:
			state.Bridge.BlockCaptain = true
		}

		if blk.Challenge != nil {
			switch blk.Challenge.Subject {
			case gm.Players[one]:
				state.Bridge.BlockChallengeSubjectOne = true
			case gm.Players[two]:
				state.Bridge.BlockChallengeSubjectTwo = true
			}
		}
	}

	fmt.Println(state)
	fmt.Println(state.Bridge)

	if !second {
		next := 0
		for i, player := range gm.Players[1:] {
			if player.Alive() {
				next = i
				break
			}
		}

		valid := []MoveEnum{}
		if gm.Players[next].Coins < 10 {
			valid = append(valid, Income)
			valid = append(valid, ForeignAid)
			valid = append(valid, Tax)
			valid = append(valid, Steal)
			valid = append(valid, Exchange)
		}

		if gm.Players[next].Coins >= 7 {
			valid = append(valid, Coup)
		}

		if gm.Players[next].Coins >= 3 {
			valid = append(valid, Assassinate)
		}

		if self == gm.Players[next] {
			// mover
		} else {
			// challenger
		}
	} else {
		// is blocker

		// is challenger

	}

	a.States = append(a.States, state)
}

func Score(state *State, action *Action) float64 {
	return 0.0
}

func (a *Agent) ChooseMove(moves []*Move) *Move {
	return NewRandom().ChooseMove(moves)
}

func (a *Agent) ChooseBlock(claims []*Claim) *Claim {
	return NewRandom().ChooseBlock(claims)
}

func (a *Agent) ChooseChallenge(claim *Claim) bool {
	return NewRandom().ChooseChallenge(claim)
}

func (a *Agent) ChooseDiscard(hand *Cards, amt int) []CardEnum {
	return NewRandom().ChooseDiscard(hand, amt)
}
