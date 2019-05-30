func maxChunksToSorted(arr []int) int {
	sorted := append([]int(nil), arr...)
	sort.Ints(sorted)

	mapA := make(map[int]int)
	mapS := make(map[int]int)
	ans := 0
	for i := range arr {
		a := arr[i]
		if mapS[a] > 0 {
			mapS[a]--
			if mapS[a] == 0 {
				delete(mapS, a)
			}
		} else {
			mapA[a]++
		}

		s := sorted[i]
		if mapA[s] > 0 {
			mapA[s]--
			if mapA[s] == 0 {
				delete(mapA, s)
			}
		} else {
			mapS[s]++
		}

		if len(mapA) == 0 && len(mapS) == 0 {
			ans++
		}
	}
	return ans
}

