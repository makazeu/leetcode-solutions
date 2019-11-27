func minTimeToVisitAllPoints(points [][]int) int {
	sx, sy := points[0][0], points[0][1]
	ans := 0
	for i := 1; i < len(points); i++ {
		ans += max(abs(points[i][0]-sx), abs(points[i][1]-sy))
		sx, sy = points[i][0], points[i][1]
	}
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

