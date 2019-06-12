func smallestSubsequence(text string) string {
	lastPos := make([]int, 26)
	for i := range lastPos {
		lastPos[i] = -1
	}
	for i := len(text) - 1; i >= 0; i-- {
		c := int(text[i] - 'a')
		if lastPos[c] == -1 {
			lastPos[c] = i
		}
	}

	var ans []rune
	used := make([]bool, 26)
	last := -1
	for {
		var head int
		for head = last + 1; head < len(text); head++ {
			c := int(text[head] - 'a')
			if used[c] {
				continue
			}
			if lastPos[c] != head {
				continue
			}
			break
		}
		if head == len(text) {
			break
		}
		var minChar uint8
		minPos := -1
		for i := last + 1; i <= head; i++ {
			c := int(text[i] - 'a')
			if used[c] {
				continue
			}
			if minPos == -1 {
				minPos = i
				minChar = text[i]
				continue
			}
			if text[i] >= minChar {
				continue
			}
			minPos = i
			minChar = text[i]
		}
		ans = append(ans, rune(minChar))
		used[int(minChar-'a')] = true
		last = minPos
	}

	return string(ans)
}

