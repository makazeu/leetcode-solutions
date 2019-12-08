func subtractProductAndSum(n int) int {
	prod := 1
	sum := 0
	for n > 0 {
		k := n % 10
		prod *= k
		sum += k
		n /= 10
	}
	return prod - sum
}

