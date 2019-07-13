func numberOfDays(Y int, M int) int {
	if M == 1 || M == 3 || M == 5 || M == 7 || M == 8 || M == 10 || M == 12 {
		return 31
	}
	if M == 4 || M == 6 || M == 9 || M == 11 {
		return 30
	}
	if (Y%100 != 0 && Y%4 == 0) || (Y%100 == 0 && Y%400 == 0) {
		return 29
	}
	return 28
}

