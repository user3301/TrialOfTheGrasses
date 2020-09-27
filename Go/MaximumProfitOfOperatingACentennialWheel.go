package gosoln

func MinOperationsMaxProfit(customers []int, boardingCost, runningCost int) int {
	currentWaiting := 0
	for _, v := range customers {
		currentWaiting += v
	}
	maxProfit, currentProfit, rotated, bestRotated := 0, 0, 0, 0
	for currentWaiting > 0 {
		if currentWaiting >= 4 {
			currentProfit += boardingCost*4 - runningCost
			currentWaiting -= 4
		} else {
			currentProfit += boardingCost*currentWaiting - runningCost
			currentWaiting = 0
		}
		rotated++
		if maxProfit < currentProfit {
			maxProfit = currentProfit
			bestRotated = rotated
		}
	}
	if maxProfit > 0 {
		return bestRotated
	} else {
		return -1
	}
}
