func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func calc(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func dfs(dists *[][]int, pos, dist int, ans *int, flags []bool) {
	if dist >= *ans {
		return
	}
	if pos == len(*dists) {
		*ans = dist
		return
	}

	for i, d := range (*dists)[pos] {
		if flags[i] {
			continue
		}
		flags[i] = true
		dfs(dists, pos+1, dist+d, ans, flags)
		flags[i] = false
	}
}

func assignBikes(workers [][]int, bikes [][]int) int {
	dists := make([][]int, len(workers))
	for i, worker := range workers {
		for _, bike := range bikes {
			dists[i] = append(dists[i], calc(worker, bike))
		}
	}
	ans := math.MaxInt32
	dfs(&dists, 0, 0, &ans, make([]bool, len(bikes)))
	return ans
}

