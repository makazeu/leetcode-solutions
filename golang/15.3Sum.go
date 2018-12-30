const NUM = 3

var pool map[int]int

func threeSum(nums []int) [][]int {
	var answers [][]int
	pool = make(map[int]int)

	sort.Ints(nums)
	answer := make([]int, NUM)
	for _, num := range nums {
		t, exists := pool[num]
		if exists {
			pool[num] = t + 1
		} else {
			pool[num] = 1
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		pool[nums[i]] = pool[nums[i]] - 1
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			pool[nums[j]] = pool[nums[j]] - 1
			ret := 0 - nums[i] - nums[j]
			if ret >= nums[j] && pool[ret] > 0 {
				answer[0] = nums[i]
				answer[1] = nums[j]
				answer[2] = ret

				result := make([]int, NUM)
				copy(result, answer)
				answers = append(answers, result)
			}
			pool[nums[j]] = pool[nums[j]] + 1
		}
		pool[nums[i]] = pool[nums[i]] + 1
	}

	return answers
}
