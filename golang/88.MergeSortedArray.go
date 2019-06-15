func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if m >= 1 && n >= 1 {
			if nums1[m-1] >= nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
			continue
		}
		if m > 0 {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}

