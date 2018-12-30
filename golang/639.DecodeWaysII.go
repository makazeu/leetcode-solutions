const MOD int64 = 1000000007

func dfs(nums []int, start, end int) int64 {
	if start > end {
		return 1
	}
	if nums[start] == 0 {
		return 0
	}
	if start == end {
		if nums[start] != -1 {
			return 1
		} else {
			return 9
		}
	}

	mid := (start + end) / 2

	num := int64(0)
	pre := dfs(nums, start, mid)
	post := dfs(nums, mid+1, end)
	num += (pre * post) % MOD

	factor := int64(1)
	if nums[mid] != -1 && nums[mid+1] != -1 {
		value := nums[mid]*10 + nums[mid+1]
		if value < 10 || value > 26 {
			return num
		}
	} else if nums[mid] == -1 && nums[mid+1] != -1 {
		if nums[mid+1] <= 6 {
			factor = 2
		}
	} else if nums[mid] != -1 && nums[mid+1] == -1 {
		switch nums[mid] {
		case 1:
			factor = 9
			break
		case 2:
			factor = 6
			break
		default:
			return num
		}
	} else if nums[mid] == -1 && nums[mid+1] == -1 {
		factor = 15
	}

	pre = dfs(nums, start, mid-1)
	post = dfs(nums, mid+2, end)
	num2 := ((factor * pre) % MOD * post) % MOD

	return (num + num2) % MOD
}

func numDecodings(s string) int {
	nums := make([]int, len(s))
	for i, c := range s {
		if c == '*' {
			nums[i] = -1
		} else {
			nums[i] = int(c - '0')
		}
	}

	return int(dfs(nums, 0, len(nums)-1))
}
