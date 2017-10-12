package coup

import (
	"math/rand"
	"time"
)

type Agent struct {
	States []*State
	Action *Action
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Update(self *Player, gm *Game, mv *Move, blk *Block, second bool) {

	state := NewState(self, gm, mv, blk)
	a.States = append(a.States, state)

	// fmt.Printf("%+v\n", state)
	// fmt.Printf("%+v\n", state.Bridge)

	actions := []*Action{}

	if !second {
		actions = MoveAndChallenges(self, gm)
	} else {
		actions = BlockAndChallenges()
	}

	// for _, action := range actions {
	// 	fmt.Printf("%+v\n", action.Move)
	// }

	bestScore := -1.0
	bestActions := []*Action{&Action{}}
	for _, action := range actions {
		score := Score(state, action)
		if score > bestScore {
			bestScore = score
			bestActions = []*Action{action}
		} else if score == bestScore {
			bestActions = append(bestActions, action)
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	a.Action = bestActions[rand.Intn(len(bestActions))]

}

func MoveAndChallenges(self *Player, gm *Game) []*Action {
	index := 0
	for i, player := range gm.Players {
		if self == player {
			index = i
		}
	}

	players := append(gm.Players[index:], gm.Players[:index]...)

	next := 0
	for i, player := range gm.Players[1:] {
		if player.Alive() {
			next = i
			break
		}
	}

	subject := next - index
	if subject < 0 {
		subject += len(players)
	}

	valid := []MoveEnum{}
	if players[subject].Coins < 10 {
		valid = append(valid, Income)
		valid = append(valid, ForeignAid)
		valid = append(valid, Tax)
		valid = append(valid, Steal)
		valid = append(valid, Exchange)
	}

	if players[subject].Coins >= 7 {
		valid = append(valid, Coup)
	}

	if players[subject].Coins >= 3 {
		valid = append(valid, Assassinate)
	}

	objects := []int{}
	for i, player := range players {
		if player.Alive() && i != subject {
			objects = append(objects, i)
		}
	}

	if subject != 0 {
		return MoveChallenges()
	} else {
		return Moves(valid, objects)
	}
}

func Moves(valid []MoveEnum, objects []int) []*Action {
	actions := []*Action{}
	for _, move := range valid {
		switch move {
		case Income:
			action := &Action{
				Move: &ActionMove{
					Income: true,
				},
			}
			actions = append(actions, action)
		case ForeignAid:
			action := &Action{
				Move: &ActionMove{
					ForeignAid: true,
				},
			}
			actions = append(actions, action)
		case Coup:
			for _, object := range objects {
				action := &Action{
					Move: &ActionMove{
						Coup: true,
					},
				}
				switch object {
				case 0:
					action.Move.ObjectOne = true
				case 1:
					action.Move.ObjectTwo = true
				case 2:
					action.Move.ObjectThree = true
				case 3:
					action.Move.ObjectFour = true
				case 4:
					action.Move.ObjectFive = true
				}
				actions = append(actions, action)
			}
		case Tax:
			for _, challengeable := range NewChallengeables(objects) {
				action := &Action{
					Move: &ActionMove{
						Tax: challengeable,
					},
				}
				actions = append(actions, action)
			}
		case Assassinate:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects) {
					action := &Action{
						Move: &ActionMove{
							Assassinate: challengeable,
						},
					}
					switch object {
					case 0:
						action.Move.ObjectOne = true
					case 1:
						action.Move.ObjectTwo = true
					case 2:
						action.Move.ObjectThree = true
					case 3:
						action.Move.ObjectFour = true
					case 4:
						action.Move.ObjectFive = true
					}
					actions = append(actions, action)
				}
			}
		case Exchange:
			for _, challengeable := range NewChallengeables(objects) {
				action := &Action{
					Move: &ActionMove{
						Exchange: challengeable,
					},
				}
				actions = append(actions, action)
			}
		case Steal:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects) {
					action := &Action{
						Move: &ActionMove{
							Steal: challengeable,
						},
					}
					switch object {
					case 0:
						action.Move.ObjectOne = true
					case 1:
						action.Move.ObjectTwo = true
					case 2:
						action.Move.ObjectThree = true
					case 3:
						action.Move.ObjectFour = true
					case 4:
						action.Move.ObjectFive = true
					}
					actions = append(actions, action)
				}
			}
		}
	}

	return actions
}

func MoveChallenges() []*Action {
	return []*Action{}
}

func BlockAndChallenges() []*Action {
	return []*Action{}
}

func Blocks() []*Action {
	return []*Action{}
}

func BlockChallenges() []*Action {
	return []*Action{}
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
