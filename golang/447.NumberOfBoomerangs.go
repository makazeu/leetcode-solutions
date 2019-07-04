func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i, point1 := range points {
		h := make(map[float64]int)
		for j, point2 := range points {
			if i == j {
				continue
			}
			h[calcDistance(point1, point2)]++
		}
		for _, v := range h {
			ans += v * (v - 1)
		}
	}
	return ans
}

func calcDistance(a, b []int) float64 {
	x, y := b[0]-a[0], b[1]-a[1]
	return math.Sqrt(float64(x*x) + float64(y*y))
}

