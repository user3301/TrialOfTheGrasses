package gosoln

func MajorityElement(nums []int) int {
	maj, counter := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == maj {
			counter++
		} else {
			counter--
			if counter <= 0 {
				maj = nums[i]
				counter = 1
			}
		}
	}
	return maj
}
