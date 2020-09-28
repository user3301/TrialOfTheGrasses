package gosoln

import "github.com/user3301/TrialOfTheGrasses/Go/pkg/utils"

func MinOperationsMaxProfit(customers []int, boardingCost, runningCost int) int {
	curWaiting, maxProfit, curProfit, rotated, bestRotated := 0, 0, 0, 0, 0
	i := 0
	for curWaiting > 0 || i < len(customers) {
		if i < len(customers) {
			curWaiting += customers[i]
			i++
		}
		boarded := utils.MinInt(4, curWaiting)
		curProfit += boardingCost*boarded - runningCost
		curWaiting -= boarded
		rotated++
		if curProfit > maxProfit {
			maxProfit = curProfit
			bestRotated = rotated
		}
	}
	if maxProfit > 0 {
		return bestRotated
	} else {
		return -1
	}
}
