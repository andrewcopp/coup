package coup

import (
	"fmt"
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
		actions = BlockAndChallenges(gm, mv, self)
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
			next = i + 1
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
		return MoveChallenges(self, subject, valid, objects)
	} else {
		return Moves(self, valid, objects)
	}
}

func Moves(self *Player, valid []MoveEnum, objects []int) []*Action {
	actions := []*Action{}
	for _, move := range valid {
		switch move {
		case Income:
			move := NewActionMove()
			move.Income = true
			action := &Action{
				Move: move,
			}
			actions = append(actions, action)
		case ForeignAid:
			move := NewActionMove()
			move.ForeignAid = true
			action := &Action{
				Move: move,
			}
			actions = append(actions, action)
		case Coup:
			for _, object := range objects {
				move := NewActionMove()
				move.Coup = true
				action := &Action{
					Move: move,
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
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				move := NewActionMove()
				move.Tax = challengeable
				action := &Action{
					Move: move,
				}
				actions = append(actions, action)
			}
		case Assassinate:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects, self.Hand) {
					move := NewActionMove()
					move.Assassinate = challengeable
					action := &Action{
						Move: move,
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
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				move := NewActionMove()
				move.Exchange = challengeable
				action := &Action{
					Move: move,
				}
				actions = append(actions, action)
			}
		case Steal:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects, self.Hand) {
					move := NewActionMove()
					move.Steal = challengeable
					action := &Action{
						Move: move,
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

func MoveChallenges(self *Player, sub int, valid []MoveEnum, objects []int) []*Action {
	action := &Action{
		ChallengeMove: NewActionMoveChallenge(),
	}
	switch sub {
	case 0:
		action.ChallengeMove.SubjectOne = true
	case 1:
		action.ChallengeMove.SubjectTwo = true
	case 2:
		action.ChallengeMove.SubjectThree = true
	case 3:
		action.ChallengeMove.SubjectFour = true
	case 4:
		action.ChallengeMove.SubjectFive = true
	}

	actions := []*Action{action}
	for _, move := range valid {
		switch move {
		case Tax:
			temp := actions
			for _, action := range actions {
				for _, reveal := range NewReveals(self.Hand) {
					c := action.ChallengeMove.Copy()
					c.Tax = reveal
					temp = append(temp, &Action{ChallengeMove: c})
				}
			}
			actions = temp
		case Assassinate:
			for _, object := range objects {
				switch object {
				case 0:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.AssassinateObjectOne = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 1:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.AssassinateObjectTwo = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 2:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.AssassinateObjectThree = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 3:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.AssassinateObjectFour = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 4:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.AssassinateObjectFive = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				}
			}
		case Exchange:
			temp := actions
			for _, action := range actions {
				for _, reveal := range NewReveals(self.Hand) {
					c := action.ChallengeMove.Copy()
					c.Exchange = reveal
					temp = append(temp, &Action{ChallengeMove: c})
				}
			}
			actions = temp
		case Steal:
			for _, object := range objects {
				switch object {
				case 0:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.StealObjectOne = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 1:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.StealObjectTwo = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 2:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.StealObjectThree = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 3:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.StealObjectFour = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				case 4:
					temp := actions
					for _, action := range actions {
						for _, reveal := range NewReveals(self.Hand) {
							c := action.ChallengeMove.Copy()
							c.StealObjectFive = reveal
							temp = append(temp, &Action{ChallengeMove: c})
						}
					}
					actions = temp
				}
			}
		}
	}

	return actions
}

func BlockAndChallenges(gm *Game, mv *Move, self *Player) []*Action {

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
			next = i + 1
			break
		}
	}

	subject := next - index
	if subject < 0 {
		subject += len(players)
	}

	objects := []int{}
	for i, player := range players {
		if player.Alive() && i != subject {
			objects = append(objects, i)
		}
	}

	actions := []*Action{&Action{
		Block:          NewActionBlock(),
		ChallengeBlock: NewActionBlockChallenge(),
	}}
	switch mv.Case {
	case ForeignAid:
		if mv.Subject != self {
			actions = append(actions, Blocks(gm, mv, self, objects)...)
		}

		actions = append(actions, BlockChallenges(gm, mv, self, objects)...)
	case Assassinate:
		if mv.Object != self {
			actions = append(actions, BlockChallenges(gm, mv, self, objects)...)
		} else {
			actions = append(actions, Blocks(gm, mv, self, objects)...)
		}
	case Steal:
		if mv.Object != self {
			actions = append(actions, BlockChallenges(gm, mv, self, objects)...)
		} else {
			actions = append(actions, Blocks(gm, mv, self, objects)...)
		}
	}
	return actions
}

func Blocks(gm *Game, mv *Move, self *Player, objects []int) []*Action {
	actions := []*Action{&Action{
		Block: NewActionBlock(),
	}}
	switch mv.Case {
	case ForeignAid:
		for _, challengeable := range NewChallengeables(objects, self.Hand) {
			block := NewActionBlock()
			block.Challengeable = challengeable
			action := &Action{
				Block: block,
			}
			actions = append(actions, action)
		}
	case Assassinate:
		if self == mv.Object {
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				block := NewActionBlock()
				block.Challengeable = challengeable
				action := &Action{
					Block: block,
				}
				actions = append(actions, action)
			}
		}
	case Steal:
		if self == mv.Object {
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				block := NewActionBlock()
				block.Challengeable = challengeable
				block.Ambassador = true
				action := &Action{
					Block: block,
				}
				actions = append(actions, action)
			}
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				block := NewActionBlock()
				block.Challengeable = challengeable
				block.Captain = true
				action := &Action{
					Block: block,
				}
				actions = append(actions, action)
			}
		}
	}
	return actions
}

func BlockChallenges(gm *Game, mv *Move, self *Player, objects []int) []*Action {
	actions := []*Action{&Action{
		ChallengeBlock: NewActionBlockChallenge(),
	}}
	switch mv.Case {
	case ForeignAid:
		for i := range objects {
			switch i {
			case 1:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						c := action.ChallengeBlock.Copy()
						c.SubjectTwoDuke = reveal
						temp = append(temp, &Action{ChallengeBlock: c})
					}
				}
				actions = temp
			case 2:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						c := action.ChallengeBlock.Copy()
						c.SubjectThreeDuke = reveal
						temp = append(temp, &Action{ChallengeBlock: c})
					}
				}
				actions = temp
			case 3:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						c := action.ChallengeBlock.Copy()
						c.SubjectFourDuke = reveal
						temp = append(temp, &Action{ChallengeBlock: c})
					}
				}
				actions = temp
			case 4:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						c := action.ChallengeBlock.Copy()
						c.SubjectFiveDuke = reveal
						temp = append(temp, &Action{ChallengeBlock: c})
					}
				}
				actions = temp
			}
		}
	case Assassinate:
		temp := actions
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				c := action.ChallengeBlock.Copy()
				c.Contessa = reveal
				temp = append(temp, &Action{ChallengeBlock: c})
			}
		}
		actions = temp
	case Steal:
		temp := actions
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				c := action.ChallengeBlock.Copy()
				c.Ambassador = reveal
				temp = append(temp, &Action{ChallengeBlock: c})
			}
		}
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				c := action.ChallengeBlock.Copy()
				c.Captain = reveal
				temp = append(temp, &Action{ChallengeBlock: c})
			}
		}
		actions = temp
	}
	return actions
}

func Score(state *State, action *Action) float64 {
	return 0.0
}

func (a *Agent) ChooseMove(gm *Game, moves []*Move) *Move {
	for _, move := range moves {
		switch move.Case {
		case Income:
			if a.Action.Move.Income {
				return move
			}
		case ForeignAid:
			if a.Action.Move.ForeignAid {
				return move
			}
		case Coup:
			if a.Action.Move.Coup {
				if a.Action.Move.ObjectOne && (gm.Players[0] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectTwo && (gm.Players[1] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectThree && (gm.Players[2] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFour && (gm.Players[3] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFive && (gm.Players[4] == move.Object) {
					return move
				}
			}
		case Tax:
			if a.Action.Move.Tax.Selected() {
				return move
			}
		case Assassinate:
			if a.Action.Move.Assassinate.Selected() {
				if a.Action.Move.ObjectOne && (gm.Players[0] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectTwo && (gm.Players[1] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectThree && (gm.Players[2] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFour && (gm.Players[3] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFive && (gm.Players[4] == move.Object) {
					return move
				}
			}
		case Exchange:
			if a.Action.Move.Exchange.Selected() {
				return move
			}
		case Steal:
			if a.Action.Move.Steal.Selected() {
				if a.Action.Move.ObjectOne && (gm.Players[0] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectTwo && (gm.Players[1] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectThree && (gm.Players[2] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFour && (gm.Players[3] == move.Object) {
					return move
				}
				if a.Action.Move.ObjectFive && (gm.Players[4] == move.Object) {
					return move
				}
			}
		}
	}

	fmt.Println(a.Action.Move)
	panic("All moves should have been explored.")
}

func (a *Agent) ChooseBlock(claims []*Claim) *Claim {
	for _, claim := range claims {
		switch claim.Declared {
		case Duke:
			if a.Action.Block.Challengeable.Selected() {
				return claim
			}
		case Ambassador:
			if a.Action.Block.Challengeable.Selected() && a.Action.Block.Ambassador {
				return claim
			}
		case Captain:
			if a.Action.Block.Challengeable.Selected() && a.Action.Block.Captain {
				return claim
			}
		case Contessa:
			if a.Action.Block.Challengeable.Selected() {
				return claim
			}
		}
	}
	return nil
}

func (a *Agent) ChooseChallengeMove(gm *Game, self *Player, claim *Claim, object *Player) bool {
	index := 0
	for i, player := range gm.Players {
		if self == player {
			index = i
		}
	}

	players := append(gm.Players[index:], gm.Players[:index]...)

	switch claim.Declared {
	case Duke:
		return a.Action.ChallengeMove.Tax.Selected()
	case Ambassador:
		return a.Action.ChallengeMove.Exchange.Selected()
	}

	obj := 0
	for i, player := range players {
		if object == player {
			obj = i
		}
	}

	switch claim.Declared {
	case Assassin:
		switch obj {
		case 0:
			return a.Action.ChallengeMove.AssassinateObjectOne.Selected()
		case 1:
			return a.Action.ChallengeMove.AssassinateObjectTwo.Selected()
		case 2:
			return a.Action.ChallengeMove.AssassinateObjectThree.Selected()
		case 3:
			return a.Action.ChallengeMove.AssassinateObjectFour.Selected()
		case 4:
			return a.Action.ChallengeMove.AssassinateObjectFive.Selected()
		}
	case Captain:
		switch obj {
		case 0:
			return a.Action.ChallengeMove.StealObjectOne.Selected()
		case 1:
			return a.Action.ChallengeMove.StealObjectTwo.Selected()
		case 2:
			return a.Action.ChallengeMove.StealObjectThree.Selected()
		case 3:
			return a.Action.ChallengeMove.StealObjectFour.Selected()
		case 4:
			return a.Action.ChallengeMove.StealObjectFive.Selected()
		}
	}

	panic("Challenge Not Found")
}

func (a *Agent) ChooseChallengeBlock(gm *Game, self *Player, claim *Claim, object *Player) bool {
	index := 0
	for i, player := range gm.Players {
		if self == player {
			index = i
		}
	}

	players := append(gm.Players[index:], gm.Players[:index]...)

	obj := 0
	for i, player := range players {
		if claim.Subject == player {
			obj = i
		}
	}

	switch claim.Declared {
	case Duke:
		switch obj {
		case 0:
			if a.Action.ChallengeBlock.SubjectOneDuke.Selected() {
				return true
			}
		case 1:
			if a.Action.ChallengeBlock.SubjectTwoDuke.Selected() {
				return true
			}
		case 2:
			if a.Action.ChallengeBlock.SubjectThreeDuke.Selected() {
				return true
			}
		case 3:
			if a.Action.ChallengeBlock.SubjectFourDuke.Selected() {
				return true
			}
		case 4:
			if a.Action.ChallengeBlock.SubjectFiveDuke.Selected() {
				return true
			}
		}
	case Ambassador:
		if a.Action.ChallengeBlock.Ambassador.Selected() {
			return true
		}
	case Captain:
		if a.Action.ChallengeBlock.Captain.Selected() {
			return true
		}
	case Contessa:
		if a.Action.ChallengeBlock.Contessa.Selected() {
			return true
		}
	}
	return false
}

func (a *Agent) ChooseDiscard(hand *Cards, amt int) []CardEnum {
	return NewRandom().ChooseDiscard(hand, amt)
}
