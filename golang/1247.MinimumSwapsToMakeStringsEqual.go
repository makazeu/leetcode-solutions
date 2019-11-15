func minimumSwap(s1 string, s2 string) int {
	var a []int
	for i, c1 := range s1 {
		if uint8(c1) != s2[i] {
			a = append(a, int(c1-'x'))
		}
	}
	fmt.Println(a)
	if len(a) == 0 {
		return 0
	}

	left := make([]int, 2)
	for _, x := range a {
		left[x]++
	}

	fixed := make([]int, 2)
	ans := 0
	for _, x := range a {
		if fixed[x] > 0 {
			fixed[x]--
			continue
		}

		left[x]--
		if left[x] > 0 {
			left[x]--
			fixed[x]++
			ans++
			continue
		}

		if left[x^1] > 0 {
			fixed[x^1]++
			left[x^1]--
			ans += 2
			continue
		}
		return -1
	}
	return ans
}

