func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	if intervals == nil || len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}

	if newInterval[1] < intervals[0][0] {
		result = append([][]int{newInterval}, intervals...)
		return result
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		result = append(intervals, newInterval)
		return result
	}

	flag := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			result = append(result, interval)
			continue
		}
		if interval[0] <= newInterval[1] && interval[1] >= newInterval[0] {
			newInterval[0] = min(interval[0], newInterval[0])
			newInterval[1] = max(interval[1], newInterval[1])
		}
		if interval[0] <= newInterval[1] {
			newInterval[1] = max(newInterval[1], interval[1])
		} else if interval[0] > newInterval[1] {
			if !flag  {
				result = append(result, newInterval)
				flag = true
			}
			result = append(result, interval)
		}
	}
	if !flag {
		result = append(result, newInterval)
	}

	return result
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

