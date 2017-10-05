package agent

func Reward(state *State) float64 {
	self := state.OneDukes + state.OneAssassins + state.OneAmbassadors + state.OneCaptains + state.OneContessa
	others := state.TwoCards

	switch self + others {
	case self:
		return 1.0
	case others:
		return -1.0
	default:
		return 0.0
	}
}
