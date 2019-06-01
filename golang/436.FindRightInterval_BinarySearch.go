type Item struct {
	s, e  int
	index int
}

type Intervals []*Item

func (intervals Intervals) Len() int {
	return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
	if intervals[i].s != intervals[j].s {
		return intervals[i].s < intervals[j].s
	}
	return intervals[i].index < intervals[j].index
}

func (intervals Intervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func findRightInterval(intervals [][]int) []int {
	var inters Intervals
	for i, e := range intervals {
		inters = append(inters, &Item{e[0], e[1], i})
	}
	sort.Sort(inters)

	var ans []int
	m := make(map[int]int)
	for _, e := range intervals {
		v, exists := m[e[1]]
		if !exists {
			v = binarySearch(inters, e[1])
			m[e[1]] = v
		}
		ans = append(ans, v)
	}
	return ans
}

func binarySearch(intervals Intervals, target int) int {
	l, r := 0, len(intervals)-1
	mid, ans := 0, -1
	for l <= r {
		mid = (l + r) / 2
		if intervals[mid].s >= target {
			ans = mid
			r = mid - 1
		} else {
			l = l + 1
		}
	}
	if ans == -1 {
		return -1
	}
	return intervals[ans].index
}

