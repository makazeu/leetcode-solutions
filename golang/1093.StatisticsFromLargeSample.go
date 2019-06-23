func sampleStats(count []int) []float64 {
	num := 0
	sum := 0
	min := math.MaxInt32
	max := 0

	modeValue := 0
	modeNum := 0
	for i, c := range count {
		num += c
		sum += c * i
		if c > 0 && min == math.MaxInt32 {
			min = i
		}
		if c > 0 {
			max = i
		}
		if c > modeNum {
			modeNum = c
			modeValue = i
		}
	}

	var median float64
	if num == 1 {
		median = float64(min)
	} else if num&1 == 1 {
		median = float64(getPosNum(count, (num+1)/2))
	} else {
		median = (float64(getPosNum(count, num/2)) + float64(getPosNum(count, num/2+1))) / 2
	}

	ans := make([]float64, 5)
	ans[0] = float64(min)                // minimum
	ans[1] = float64(max)                // maximum
	ans[2] = float64(sum) / float64(num) // mean
	ans[3] = median                      // median
	ans[4] = float64(modeValue)          // mode
	return ans
}

func getPosNum(count []int, pos int) int {
	n := 0
	for i, c := range count {
		if n+c >= pos {
			return i
		}
		n += c
	}
	panic("no such num")
}

