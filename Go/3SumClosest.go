package Go

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		panic("impossible!")
	}

	sort.Ints(nums)

	closeset := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum > target:
				for r-1 > l && nums[r-1] == nums[r] {
					r--
				}
				r--
			case sum < target:
				for l+1 < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			default:
				return sum
			}
			if Abs(target-sum) < Abs(target-closeset) {
				closeset = sum
			}
		}
	}
	return closeset
}
