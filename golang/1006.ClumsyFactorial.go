func clumsy(N int) int {
	if N <= 3 {
		return initials(N)
	}

	var ans = calcFirst(N)
	N -= 4
	for N >= 4 {
		ans -= calc(N)
		N -= 4
	}

	if N > 0 {
		ans -= int64(initials(N))
	}

	return int(ans)
}

func calcFirst(n int) int64 {
	return int64(n)*int64(n-1)/int64(n-2) + int64(n-3)
}

func calc(n int) int64 {
	return int64(n)*int64(n-1)/int64(n-2) - int64(n-3)
}

func initials(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return 6
}

