func prevPermOpt1(A []int) []int {
	if len(A) < 2 {
		return A
	}
	pos := -1
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			pos = i
			break
		}
	}
	if pos == -1 {
		return A
	}

	maxV, maxP := 0, 0
	for i := len(A) - 1; i > pos; i-- {
		if A[i] != A[pos] && A[i] < A[pos] && A[i] > maxV {
			maxV = A[i]
			maxP = i
		}
	}

	A[pos], A[maxP] = A[maxP], A[pos]
	return A
}

