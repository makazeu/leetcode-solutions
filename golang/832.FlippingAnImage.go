func flipAndInvertImage(A [][]int) [][]int {
	for _, e := range A {
		l := len(e)
		for i := 0; i < (l+1)/2; i++ {
			e[i], e[l-1-i] = e[l-1-i], e[i]
		}
		for i := range e {
			e[i] ^= 1
		}
	}
	return A
}

