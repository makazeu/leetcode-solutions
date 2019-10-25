func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	k := float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])

	for i:=2; i< len(coordinates); i++ {
		kk := float64(coordinates[i][1]-coordinates[0][1]) / float64(coordinates[i][0]-coordinates[0][0])
		if kk != k {
			return false
		}
	}
	return true
}

