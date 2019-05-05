func isBoomerang(points [][]int) bool {
	if len(points) <= 1 {
		return true
	}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				return false
			}
		}
	}

	f := float64(points[1][1]-points[0][1]) / float64(points[1][0]-points[0][0])
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if j != 1 && float64(points[j][1]-points[i][1]) / float64(points[j][0]-points[i][0]) == f {
				return false
			}
		}
	}

	return true
}

