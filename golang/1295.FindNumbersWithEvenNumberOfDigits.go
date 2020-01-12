func findNumbers(nums []int) int {
	res := 0
        for _, num := range nums {
		if getLength(num) % 2 == 0 {
			res++
		}
	}
	return res
}

func getLength(num int) int {
	return int(math.Log10(float64(num))) + 1
}

