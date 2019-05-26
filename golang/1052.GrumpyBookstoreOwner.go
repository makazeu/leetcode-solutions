func maxSatisfied(customers []int, grumpy []int, X int) int {
	maxV, maxP := 0, 0
	curV := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			curV += customers[i]
		}
	}
	maxV = curV
	for i := 1; i <= len(customers)-X; i++ {
		if grumpy[i-1] == 1 {
			curV -= customers[i-1]
		}
		if grumpy[i+X-1] == 1 {
			curV += customers[i+X-1]
		}

		if curV > maxV {
			maxV = curV
			maxP = i
		}
	}

	ans := 0
	for i := range customers {
		if (i < maxP || i >= maxP+X) && grumpy[i] == 1 {
			continue
		}
		ans += customers[i]
	}
	return ans
}

