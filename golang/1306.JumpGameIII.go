func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}

	queue := list.New()
	N := len(arr)
	flag := make([]bool, N)
	queue.PushBack(start)

	for queue.Len() > 0 {
		now := queue.Front().Value.(int)
		queue.Remove(queue.Front())

		nxt := now + arr[now]
		if nxt < N && !flag[nxt] {
			if arr[nxt] == 0 {
				return true
			}
			flag[nxt] = true
			queue.PushBack(nxt)
		}
		nxt = now - arr[now]
		if nxt >= 0 && !flag[nxt] {
			if arr[nxt] == 0 {
				return true
			}
			flag[nxt] = true
			queue.PushBack(nxt)
		}
	}

	return false
}

