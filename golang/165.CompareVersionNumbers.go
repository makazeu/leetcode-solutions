func compareVersion(version1 string, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")
	l1, l2 := 0, 0

	for l1 < len(ver1) || l2 < len(ver2) {
		v1, v2 := 0, 0
		if l1 < len(ver1) {
			v1, _ = strconv.Atoi(ver1[l1])
		}
		if l2 < len(ver2) {
			v2, _ = strconv.Atoi(ver2[l2])
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		l1++
		l2++
	}

	return 0
}
