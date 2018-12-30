func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	sort.Ints(num)
	n := len(num)
	var med float64
	if n%2 == 0 {
		med = (float64(num[n/2-1]) + float64(num[n/2])) / 2
	} else {
		med = float64(num[(n+1)/2-1])
	}
	return med
}
