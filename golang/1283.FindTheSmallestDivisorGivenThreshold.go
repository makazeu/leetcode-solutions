func smallestDivisor(nums []int, threshold int) int {
	sort.Ints(nums)
	l := 1
	r := nums[len(nums)-1]
	ans := r
	for l <= r {
		mid := (l + r) / 2
		sum := calc(nums, mid)
		if sum <= threshold {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func calc(nums []int, div int) int {
	sum := 0
	for _, num := range nums {
		sum += int(math.Ceil(float64(num) / float64(div)))
	}
	return sum
}

