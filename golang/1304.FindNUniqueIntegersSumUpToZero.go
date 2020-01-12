func sumZero(n int) []int {
	var res []int
	for i := 1; i <= n/2; i++ {
		res = append(res, i)
		res = append(res, -i)
	}
	if n % 2 == 1 {
		res = append(res, 0)
	}
	return res
}

