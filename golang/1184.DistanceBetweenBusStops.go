func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	sum := 0
	for _, d := range distance {
		sum += d
	}

	sum2 := 0
	for i := start; i < destination; i++ {
		sum2 += distance[i]
	}

	if sum-sum2 < sum2 {
		return sum - sum2
	}
	return sum2
}

