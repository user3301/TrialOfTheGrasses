package Go

import "sort"

func ThreeSum(nums []int) [][]int {
	var ans [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		cur := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			switch {
			case nums[l]+nums[r]+cur == 0:
				item := []int{nums[l], nums[r], cur}
				ans = append(ans, item)
				for l+1 < r && nums[l+1] == nums[l] {
					l++
				}
				for r-1 > l && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			case nums[l]+nums[r]+cur > 0:
				for r-1 > l && nums[r-1] == nums[r] {
					r--
				}
				r--
			default:
				for l+1 < r && nums[l+1] == nums[l] {
					l++
				}
				l++
			}
		}
		for i+1 < len(nums)-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return ans
}
