func freqAlphabets(s string) string {
	strs := strings.Split(s, "#")
	strNum := len(strs)
	flag := false
	if len(strs[strNum-1]) == 0 {
		strs = strs[:strNum-1]
		flag = true
		strNum--
	}

	var ans []rune
	for p := 0; p < strNum-1; p++ {
		str := strs[p]
		l := len(str)
		for i := 0; i < l-2; i++ {
			ans = append(ans, rune(str[i]-'1')+'a')
		}

		ans = append(ans, rune((str[l-2]-'0')*10+str[l-1]-'0'-1)+'a')
	}

	if flag {
		str := strs[strNum-1]
		l := len(str)
		for i := 0; i < l-2; i++ {
			ans = append(ans, rune(str[i]-'1')+'a')
		}

		ans = append(ans, rune((str[l-2]-'0')*10+str[l-1]-'0'-1)+'a')
	} else {
		str := strs[strNum-1]
		l := len(str)
		for i := 0; i < l; i++ {
			ans = append(ans, rune(str[i]-'1')+'a')
		}
	}
	return string(ans)
}

