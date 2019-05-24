func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		s := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(s[0])
		domain := strings.Split(s[1], ".")
		for i := 0; i < len(domain); i++ {
			d := strings.Join(domain[i:], ".")
			m[d] += count
		}
	}

	var ans []string
	for k, v := range m {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}

