type PointSlice [][]int

func (pointSlice PointSlice) Len() int {
	return len(pointSlice)
}

func (pointSlice PointSlice) Less(i, j int) bool {
	a := pointSlice[i][0]*pointSlice[i][0] + pointSlice[i][1]*pointSlice[i][1]
	b := pointSlice[j][0]*pointSlice[j][0] + pointSlice[j][1]*pointSlice[j][1]
	return a < b
}

func (pointSlice PointSlice) Swap(i, j int) {
	pointSlice[i], pointSlice[j] = pointSlice[j], pointSlice[i]
}

func kClosest(points [][]int, K int) [][]int {
	sort.Sort(PointSlice(points))
	return points[:K]
}

