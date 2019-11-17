func encode(num int) string {
	return strconv.FormatInt(int64(num)+1, 2)[1:]
}

