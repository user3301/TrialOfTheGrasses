package gosoln

func SearchInRotatedSortedArray(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	shuffle := l

	l, r = 0, len(nums)-1

	for l <= r {
		mid := (l + r) >> 1
		realmid := (mid + shuffle) % len(nums)
		switch {
		case nums[realmid] == target:
			return realmid
		case nums[realmid] > target:
			r = mid - 1
		default:
			l = mid + 1
		}
	}
	return -1
}
