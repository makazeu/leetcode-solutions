func check(num int) bool {
	x := num
	var k int
	for x > 0 {
		k = x % 10
		if k == 0 {
			return false
		}
		if num%k != 0 {
			return false
		}
		x = x / 10
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		r := check(i)
		if r {
			ans = append(ans, i)
		}
	}
	return ans
}

