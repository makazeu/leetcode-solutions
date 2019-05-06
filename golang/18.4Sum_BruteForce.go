func fourSum(nums []int, target int) [][]int {
	N := len(nums)
	sort.Ints(nums)
	hash := make(map[int]int)
	for _, e := range nums {
		hash[e]++
	}
	var ans [][]int

	sum := 0
	for i := 0; i <= N-4; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sum += nums[i]
		if target-sum < 3*nums[i] {
			sum -= nums[i]
			continue
		}
		hash[nums[i]]--

		for j := i + 1; j <= N-3; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum += nums[j]
			if target-sum < 2*nums[j] {
				sum -= nums[j]
				continue
			}
			hash[nums[j]]--

			for k := j + 1; k <= N-2; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				sum += nums[k]
				if target-sum < nums[k] {
					sum -= nums[k]
					continue
				}
				hash[nums[k]] --

				left := target - sum
				if hash[left] > 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[k], left})
				}

				hash[nums[k]] ++
				sum -= nums[k]
			}
			hash[nums[j]] ++
			sum -= nums[j]
		}
		hash[nums[i]] ++
		sum -= nums[i]
	}

	return ans
}

