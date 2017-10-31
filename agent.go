package coup

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Agent struct {
	Version int
	Epsilon float64
	States  []*State
	Actions []*Action
	Action  *Action
}

func NewAgent(ver int, e float64) *Agent {
	return &Agent{
		Version: ver,
		Epsilon: e,
	}
}

func (a *Agent) Record(win bool) {

	if win {
		tensor := append(a.States[len(a.States)-2].Tensor(), a.Actions[len(a.Actions)-2].Tensor()...)
		strs := make([]string, len(tensor))
		for i, t := range tensor {
			strs[i] = strconv.FormatFloat(t, 'f', 5, 64)
		}
		str := strings.Join(strs, ",")
		infile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
		outfile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
		if err := exec.Command("python3", "/home/ubuntu/reinforcement/train.py", infile, outfile, str, "1.0").Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		tensor := append(a.States[len(a.States)-2].Tensor(), a.Actions[len(a.Actions)-2].Tensor()...)
		strs := make([]string, len(tensor))
		for i, t := range tensor {
			strs[i] = strconv.FormatFloat(t, 'f', 5, 64)
		}
		str := strings.Join(strs, ",")
		infile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
		outfile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
		if err := exec.Command("python3", "/home/ubuntu/reinforcement/train.py", infile, outfile, str, "0.0").Run(); err != nil {
			fmt.Println(err)
		}
	}

	for i := len(a.States) - 2; i >= 0; i-- {
		tensor := append(a.States[i].Tensor(), a.Actions[i].Tensor()...)
		strs := make([]string, len(tensor))
		for i, t := range tensor {
			strs[i] = strconv.FormatFloat(t, 'f', 5, 64)
		}
		features := strings.Join(strs, ",")
		infile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
		outfile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)

		var labels string
		if i != len(a.States)-2 {
			tensor := append(a.States[i+1].Tensor(), a.Actions[i+1].Tensor()...)
			strs := make([]string, len(tensor))
			for i, t := range tensor {
				strs[i] = strconv.FormatFloat(t, 'f', 5, 64)
			}
			str := strings.Join(strs, ",")
			bytes, err := exec.Command("python3", "/home/ubuntu/reinforcement/fit.py", infile, str).CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}

			float, err := strconv.ParseFloat(string(bytes), 64)
			if err != nil {
				fmt.Println(err)
			}

			labels = strconv.FormatFloat(float, 'f', 5, 64)
		} else {
			if win {
				labels = "0.9"
			} else {
				labels = "0.0"
			}
		}

		if err := exec.Command("python3", "/home/ubuntu/reinforcement/train.py", infile, outfile, features, labels).Run(); err != nil {
			fmt.Println(err)
		}
	}
}

func (a *Agent) Update(self *Player, gm *Game, mv *Move, blk *Block, second bool) {

	state := NewState(self, gm, mv, blk)
	a.States = append(a.States, state)

	// fmt.Printf("%+v\n", state)
	// fmt.Printf("%+v\n", state.Bridge)

	actions := []*Action{}

	if self.Hand.Size() > 2 {
		// TODO: Choose Discard
	}

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

	if !second {
		actions = DiscardsMoveAndChallenges(players, subject, objects)
	} else {
		actions = BlockAndChallenges(gm, mv, self)
	}

	// for _, action := range actions {
	// 	fmt.Printf("%+v\n", action.Move)
	// }

	rand.Seed(time.Now().UnixNano())
	if 1.0-a.Epsilon > rand.Float64() {
		bestScore := -1.0
		bestActions := []*Action{NewAction()}
		for _, action := range actions {
			score := a.Score(state, action)
			if score > bestScore {
				bestScore = score
				bestActions = []*Action{action}
			} else if score == bestScore {
				bestActions = append(bestActions, action)
			}
		}

		rand.Seed(time.Now().UnixNano())
		action := bestActions[rand.Intn(len(bestActions))]
		a.Action = action
		a.Actions = append(a.Actions, action)
	} else {
		// TODO: Why need this?!
		if len(actions) != 0 {
			rand.Seed(time.Now().UnixNano())
			action := actions[rand.Intn(len(actions))]
			a.Action = action
			a.Actions = append(a.Actions, action)
		} else {
			action := NewAction()
			a.Action = action
			a.Actions = append(a.Actions, action)
		}
	}

}

func DiscardsMoveAndChallenges(players []*Player, sub int, objs []int) []*Action {
	valid := []MoveEnum{}
	if players[sub].Coins < 10 {
		valid = append(valid, Income)
		valid = append(valid, ForeignAid)
		valid = append(valid, Tax)
		valid = append(valid, Steal)
		valid = append(valid, Exchange)
	}

	if players[sub].Coins >= 7 {
		valid = append(valid, Coup)
	}

	if players[sub].Coins >= 3 {
		valid = append(valid, Assassinate)
	}

	hand := players[0].Hand

	if sub != 0 {
		return MoveChallenges(hand, sub, valid, objs)
	}

	return Moves(hand, valid, objs)
}

func Moves(hand *Cards, valid []MoveEnum, objects []int) []*Action {
	actions := []*Action{}
	for _, move := range valid {
		switch move {
		case Income:
			action := NewAction()
			action.Move.Income = true
			actions = append(actions, action)
		case ForeignAid:
			action := NewAction()
			action.Move.ForeignAid = true
			actions = append(actions, action)
		case Coup:
			for _, object := range objects {
				action := NewAction()
				action.Move.Coup = true
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
			for _, challengeable := range NewChallengeables(objects, hand) {
				action := NewAction()
				action.Move.Tax = challengeable
				actions = append(actions, action)
			}
		case Assassinate:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects, hand) {
					action := NewAction()
					action.Move.Assassinate = challengeable
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
			for _, challengeable := range NewChallengeables(objects, hand) {
				action := NewAction()
				action.Move.Exchange = challengeable
				actions = append(actions, action)
			}
		case Steal:
			for _, object := range objects {
				for _, challengeable := range NewChallengeables(objects, hand) {
					action := NewAction()
					action.Move.Steal = challengeable
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

func MoveChallenges(hand *Cards, sub int, valid []MoveEnum, objects []int) []*Action {

	result := []*Action{}

	for _, action := range Discards(hand) {
		action.ChallengeMove = NewActionMoveChallenge()
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
					for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
						act := action.Copy()
						act.ChallengeMove.Tax = reveal
						temp = append(temp, act)
					}
				}
				actions = temp
			case Assassinate:
				for _, object := range objects {
					switch object {
					case 0:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.AssassinateObjectOne = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 1:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.AssassinateObjectTwo = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 2:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.AssassinateObjectThree = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 3:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.AssassinateObjectFour = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 4:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.AssassinateObjectFive = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					}
				}
			case Exchange:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
						act := action.Copy()
						act.ChallengeMove.Exchange = reveal
						temp = append(temp, act)
					}
				}
				actions = temp
			case Steal:
				for _, object := range objects {
					switch object {
					case 0:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.StealObjectOne = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 1:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.StealObjectTwo = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 2:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.StealObjectThree = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 3:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.StealObjectFour = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					case 4:
						temp := actions
						for _, action := range actions {
							for _, reveal := range NewReveals(Hand(hand, action.Discard)) {
								act := action.Copy()
								act.ChallengeMove.StealObjectFive = reveal
								temp = append(temp, act)
							}
						}
						actions = temp
					}
				}
			}
		}

		result = append(result, actions...)
	}

	return result
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

	actions := []*Action{NewAction()}

	switch mv.Case {
	case ForeignAid:
		if mv.Subject != self {
			actions = append(actions, Blocks(gm, mv, self, objects, actions)...)
		}

		actions = append(actions, BlockChallenges(gm, mv, self, objects, actions)...)
	case Assassinate:
		if mv.Object != self {
			actions = append(actions, BlockChallenges(gm, mv, self, objects, actions)...)
		} else {
			actions = append(actions, Blocks(gm, mv, self, objects, actions)...)
		}
	case Steal:
		if mv.Object != self {
			actions = append(actions, BlockChallenges(gm, mv, self, objects, actions)...)
		} else {
			actions = append(actions, Blocks(gm, mv, self, objects, actions)...)
		}
	}

	return actions
}

func Blocks(gm *Game, mv *Move, self *Player, objects []int, actions []*Action) []*Action {
	temp := actions
	for _, a := range temp {
		switch mv.Case {
		case ForeignAid:
			for _, challengeable := range NewChallengeables(objects, self.Hand) {
				action := a.Copy()
				block := NewActionBlock()
				block.Challengeable = challengeable
				action.Block = block
				actions = append(actions, action)
			}
		case Assassinate:
			if self == mv.Object {
				for _, challengeable := range NewChallengeables(objects, self.Hand) {
					action := a.Copy()
					block := NewActionBlock()
					block.Challengeable = challengeable
					action.Block = block
					actions = append(actions, action)
				}
			}
		case Steal:
			if self == mv.Object {
				for _, challengeable := range NewChallengeables(objects, self.Hand) {
					action := a.Copy()
					block := NewActionBlock()
					block.Challengeable = challengeable
					block.Ambassador = true
					action.Block = block
					actions = append(actions, action)
				}
				for _, challengeable := range NewChallengeables(objects, self.Hand) {
					action := a.Copy()
					block := NewActionBlock()
					block.Challengeable = challengeable
					block.Captain = true
					action.Block = block
					actions = append(actions, action)
				}
			}
		}
	}

	return actions
}

func BlockChallenges(gm *Game, mv *Move, self *Player, objects []int, actions []*Action) []*Action {
	switch mv.Case {
	case ForeignAid:
		for i := range objects {
			switch i {
			case 1:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						a := action.Copy()
						a.ChallengeBlock.SubjectTwoDuke = reveal
						temp = append(temp, a)
					}
				}
				actions = temp
			case 2:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						a := action.Copy()
						a.ChallengeBlock.SubjectThreeDuke = reveal
						temp = append(temp, a)
					}
				}
				actions = temp
			case 3:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						a := action.Copy()
						a.ChallengeBlock.SubjectFourDuke = reveal
						temp = append(temp, a)
					}
				}
				actions = temp
			case 4:
				temp := actions
				for _, action := range actions {
					for _, reveal := range NewReveals(self.Hand) {
						a := action.Copy()
						a.ChallengeBlock.SubjectFiveDuke = reveal
						temp = append(temp, a)
					}
				}
				actions = temp
			}
		}
	case Assassinate:
		temp := actions
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				a := action.Copy()
				a.ChallengeBlock.Contessa = reveal
				temp = append(temp, a)
			}
		}
		actions = temp
	case Steal:
		temp := actions
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				a := action.Copy()
				a.ChallengeBlock.Ambassador = reveal
				temp = append(temp, a)
			}
		}
		for _, action := range actions {
			for _, reveal := range NewReveals(self.Hand) {
				a := action.Copy()
				a.ChallengeBlock.Captain = reveal
				temp = append(temp, a)
			}
		}
		actions = temp
	}
	return actions
}

func Discards(hand *Cards) []*Action {
	if hand.Size() > 2 {
		actions := []*Action{}

		if hand.Dukes > 1 {
			action := NewAction()
			action.Discard.TwoDukes = true
			actions = append(actions, action)
		}

		if hand.Assassins > 1 {
			action := NewAction()
			action.Discard.TwoAssassins = true
			actions = append(actions, action)
		}

		if hand.Ambassadors > 1 {
			action := NewAction()
			action.Discard.TwoAmbassadors = true
			actions = append(actions, action)
		}

		if hand.Captains > 1 {
			action := NewAction()
			action.Discard.TwoCaptains = true
			actions = append(actions, action)
		}

		if hand.Contessas > 1 {
			action := NewAction()
			action.Discard.TwoContessas = true
			actions = append(actions, action)
		}

		if hand.Dukes > 0 && hand.Assassins > 0 {
			action := NewAction()
			action.Discard.OneDuke = true
			action.Discard.OneAssassin = true
			actions = append(actions, action)
		}

		if hand.Dukes > 0 && hand.Ambassadors > 0 {
			action := NewAction()
			action.Discard.OneDuke = true
			action.Discard.OneAmbassador = true
			actions = append(actions, action)
		}

		if hand.Dukes > 0 && hand.Captains > 0 {
			action := NewAction()
			action.Discard.OneDuke = true
			action.Discard.OneCaptain = true
			actions = append(actions, action)
		}

		if hand.Dukes > 0 && hand.Contessas > 0 {
			action := NewAction()
			action.Discard.OneDuke = true
			action.Discard.OneContessa = true
			actions = append(actions, action)
		}

		if hand.Assassins > 0 && hand.Ambassadors > 0 {
			action := NewAction()
			action.Discard.OneAssassin = true
			action.Discard.OneAmbassador = true
			actions = append(actions, action)
		}

		if hand.Assassins > 0 && hand.Captains > 0 {
			action := NewAction()
			action.Discard.OneAssassin = true
			action.Discard.OneCaptain = true
			actions = append(actions, action)
		}

		if hand.Assassins > 0 && hand.Contessas > 0 {
			action := NewAction()
			action.Discard.OneAssassin = true
			action.Discard.OneContessa = true
			actions = append(actions, action)
		}

		if hand.Ambassadors > 0 && hand.Captains > 0 {
			action := NewAction()
			action.Discard.OneAmbassador = true
			action.Discard.OneCaptain = true
			actions = append(actions, action)
		}

		if hand.Ambassadors > 0 && hand.Contessas > 0 {
			action := NewAction()
			action.Discard.OneAmbassador = true
			action.Discard.OneContessa = true
			actions = append(actions, action)
		}

		if hand.Captains > 0 && hand.Contessas > 0 {
			action := NewAction()
			action.Discard.OneCaptain = true
			action.Discard.OneContessa = true
			actions = append(actions, action)
		}

		return actions
	}

	return []*Action{NewAction()}
}

func Hand(hand *Cards, discard *Discard) *Cards {
	hand = hand.Copy()

	if discard.OneDuke {
		hand.Remove(Duke)
	}

	if discard.TwoDukes {
		hand.Remove(Duke)
		hand.Remove(Duke)
	}

	if discard.OneAssassin {
		hand.Remove(Assassin)
	}

	if discard.TwoAssassins {
		hand.Remove(Assassin)
		hand.Remove(Assassin)
	}

	if discard.OneAmbassador {
		hand.Remove(Ambassador)
	}

	if discard.TwoAmbassadors {
		hand.Remove(Ambassador)
		hand.Remove(Ambassador)
	}

	if discard.OneCaptain {
		hand.Remove(Captain)
	}

	if discard.TwoCaptains {
		hand.Remove(Captain)
		hand.Remove(Captain)
	}

	if discard.OneContessa {
		hand.Remove(Contessa)
	}

	if discard.TwoContessas {
		hand.Remove(Contessa)
		hand.Remove(Contessa)
	}

	return hand
}

func (a *Agent) Score(state *State, action *Action) float64 {
	tensor := append(state.Tensor(), action.Tensor()...)
	strs := make([]string, len(tensor))
	for i, t := range tensor {
		strs[i] = strconv.FormatFloat(t, 'f', 5, 64)
	}
	str := strings.Join(strs, ",")
	infile := fmt.Sprintf("./cmd/trainer/models/model_%d.cptk", a.Version+1)
	bytes, err := exec.Command("python3", "/home/ubuntu/reinforcement/fit.py", infile, str).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	float, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		fmt.Println(err)
	}

	return float
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
	cards := []CardEnum{}

	if a.Action.Discard.OneDuke {
		cards = append(cards, Duke)
	}

	if a.Action.Discard.TwoDukes {
		cards = append(cards, Duke)
		cards = append(cards, Duke)
	}

	if a.Action.Discard.OneAssassin {
		cards = append(cards, Assassin)
	}

	if a.Action.Discard.TwoAssassins {
		cards = append(cards, Assassin)
		cards = append(cards, Assassin)
	}

	if a.Action.Discard.OneAmbassador {
		cards = append(cards, Ambassador)
	}

	if a.Action.Discard.TwoAmbassadors {
		cards = append(cards, Ambassador)
		cards = append(cards, Ambassador)
	}

	if a.Action.Discard.OneCaptain {
		cards = append(cards, Captain)
	}

	if a.Action.Discard.TwoCaptains {
		cards = append(cards, Captain)
		cards = append(cards, Captain)
	}

	if a.Action.Discard.OneAmbassador {
		cards = append(cards, Contessa)
	}

	if a.Action.Discard.TwoAmbassadors {
		cards = append(cards, Contessa)
		cards = append(cards, Contessa)
	}

	if len(cards) != amt {
		return NewRandom().ChooseDiscard(hand, amt)
	}

	return cards
}
