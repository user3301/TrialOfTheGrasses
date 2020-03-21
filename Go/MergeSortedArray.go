package Go

func Merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := m + n - 1

	p1, p2 := m-1, n-1

	for p1 >= 0 && p2 >= 0 {
		nums1[sorted] = func() int {
			if nums1[p1] > nums2[p2] {
				ret := nums1[p1]
				p1--
				return ret
			} else {
				ret := nums2[p2]
				p2--
				return ret
			}
		}()
		sorted--
	}

	for p2 >= 0 {
		nums1[sorted] = nums2[p2]
		sorted--
		p2--
	}
}
