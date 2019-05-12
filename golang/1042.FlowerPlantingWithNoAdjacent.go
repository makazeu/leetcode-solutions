const FlowerTypeNum = 4

func gardenNoAdj(N int, paths [][]int) []int {
	graph := make([][]int, N)
	for _, e := range paths {
		e[0]--
		e[1]--
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	flower := make([]int, N)

	for i := 0; i < N; i++ {
		flag := make([]bool, FlowerTypeNum+1)
		for _, j := range graph[i] {
			if flower[j] != 0 {
				flag[flower[j]] = true
			}
		}

		for f := 1; f <= FlowerTypeNum; f++ {
			if !flag[f] {
				flower[i] = f
				break
			}
		}
	}

	return flower
}

