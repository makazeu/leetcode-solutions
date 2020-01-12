func findBestValue(arr []int, target int) int {
	l, r := 1, target
	var mid, ans int
	for l <= r {
		mid = (l + r) / 2
		sum := calc(arr, mid)
		if sum == target {
			return mid
		}
		if sum < target {
			ans = mid
			l = l + 1
		} else {
			r = r - 1
		}
	}
	sum1 := calc(arr, ans)
	sum2 := calc(arr, ans+1)
	if target-sum1 <= sum2-target {
		return ans
	} else {
		return ans + 1
	}
}

func calc(arr []int, value int) int {
	sum := 0
	for _, x := range arr {
		if x >= value {
			sum += value
		} else {
			sum += x
		}
	}
	return sum
}

