package domain

func GetWinnerScores(isPraiser bool, scores byte) byte {
	if isPraiser {
		if scores == 120 {
			return 3
		}
		if scores > 90 {
			return 2
		}
		return 1
	}
	if scores > 90 {
		return 3
	}
	return 2
}
