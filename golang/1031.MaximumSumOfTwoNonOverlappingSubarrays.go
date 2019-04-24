func maxSumTwoNoOverlap(A []int, L int, M int) int {
	maxRightSumL := calMaxRightSumSubArray(A, L)
	maxRightSumM := calMaxRightSumSubArray(A, M)
	maxLeftSumL := calMaxLeftSumSubArray(A, L)
	maxLeftSumM := calMaxLeftSumSubArray(A, M)

	maxSum := 0
	for i := 0; i < len(A); i++ {
		if maxLeftSumL[i]+maxRightSumM[i+1] > maxSum {
			maxSum = maxLeftSumL[i] + maxRightSumM[i+1]
		}
		if maxLeftSumM[i]+maxRightSumL[i+1] > maxSum {
			maxSum = maxLeftSumM[i] + maxRightSumL[i+1]
		}
	}

	return maxSum
}

func calMaxLeftSumSubArray(A []int, length int) []int {
	l := len(A)
	maxSum := make([]int, l+1)
	if length == 0 || l < length {
		return maxSum
	}
	sum := 0
	for i := 0; i < length; i++ {
		sum += A[i]
	}
	maxSum[length-1] = sum

	for i := length; i < l; i++ {
		maxSum[i] = maxSum[i-1]
		sum = sum + A[i] - A[i-length]
		if sum > maxSum[i] {
			maxSum[i] = sum
		}
	}

	return maxSum
}

func calMaxRightSumSubArray(A []int, length int) []int {
	l := len(A)
	maxSum := make([]int, l+1)
	if length == 0 || l < length {
		return maxSum
	}
	sum := 0
	for i := l - length; i < l; i++ {
		sum += A[i]
	}
	maxSum[l-length] = sum

	for i := l - length - 1; i >= 0; i-- {
		maxSum[i] = maxSum[i+1]
		sum = sum + A[i] - A[i+length]
		if sum > maxSum[i] {
			maxSum[i] = sum
		}
	}

	return maxSum
}

