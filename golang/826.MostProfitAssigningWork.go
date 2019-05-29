type Job struct {
	difficulty int
	profit     int
}

type Jobs []*Job

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Less(i, j int) bool {
	return jobs[i].difficulty < jobs[j].difficulty
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var jobs Jobs
	for i := 0; i < len(difficulty); i++ {
		jobs = append(jobs, &Job{difficulty[i], profit[i]})
	}
	sort.Sort(jobs)

	p := 0
	for _, job := range jobs {
		p = max(p, job.profit)
		job.profit = p
	}

	ans := 0
	for _, w := range worker {
		ans += binarySearch(jobs, w)
	}

	return ans
}

func binarySearch(jobs Jobs, difficulty int) int {
	if difficulty < jobs[0].difficulty {
		return 0
	}
	ans := 0
	l, r := 0, jobs.Len()-1
	for l <= r {
		mid := (l + r) / 2
		if jobs[mid].difficulty <= difficulty {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return jobs[ans].profit
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

