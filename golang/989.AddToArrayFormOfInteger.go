func addToArrayForm(A []int, K int) []int {
	reverseIntSlice(A)

	for i := 0; ; i++ {
		k := K % 10
		K = K / 10
		a := (k + A[i]) % 10
		c := (k + A[i]) / 10

		A[i] = a
		if i+1 == len(A) {
			A = append(A, c)
		} else {
			A[i+1] += c
		}

		if K == 0 && c == 0 {
			break
		}
	}

	reverseIntSlice(A)
	if len(A) > 1 && A[0] == 0 {
		A = A[1:]
	}
	return A
}

func reverseIntSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
