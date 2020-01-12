func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	N := len(friends)
	dist := make([]int, N)
	flag := make([]bool, N)
	for i := range dist {
		dist[i] = math.MaxUint32
	}
	dist[id] = 0

	queue := list.New()
	queue.PushBack(id)
	for queue.Len() > 0 {
		now := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		flag[now] = false

		for _, nxt := range friends[now] {
			if dist[now]+1 >= dist[nxt] {
				continue
			}
			dist[nxt] = dist[now] + 1
			if flag[nxt] {
				continue
			}
			queue.PushBack(nxt)
			flag[nxt] = true
		}
	}

	video := make(map[string]int)
	for i := range friends {
		if dist[i] != level {
			continue
		}
		for _, v := range watchedVideos[i] {
			video[v]++
		}
	}

	var ans []string
	for len(video) > 0 {
		maxN := math.MaxUint32
		var maxV string

		for k, v := range video {
			if v < maxN {
				maxN = v
				maxV = k
				continue
			}
			if v > maxN {
				continue
			}
			if k > maxV {
				continue
			}
			maxN = v
			maxV = k
		}
		ans = append(ans, maxV)
		delete(video, maxV)
	}
	return ans
}

