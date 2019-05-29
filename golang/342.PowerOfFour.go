func isPowerOfFour(n int) bool {
	num := math.Log(float64(n)) / math.Log(float64(4))
	return math.Abs(num - math.Round(num)) < 1e-12
}

