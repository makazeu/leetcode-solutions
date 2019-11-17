func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	parent := make(map[string]string)
	for _, region := range regions {
		f := region[0]
		for i := 1; i < len(region); i++ {
			parent[region[i]] = f
		}
	}
	flag := make(map[string]bool)
	r := region1
	for r != "" {
		flag[r] = true
		r = parent[r]
	}

	r = region2
	for {
		if flag[r] {
			return r
		}
		r = parent[r]
	}
}

