func isPowerOfThree(n int) bool {
	num := math.Log(float64(n)) / math.Log(float64(3))
	return math.Abs(num - math.Round(num)) < 1e-12
}

