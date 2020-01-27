func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    dist := make([][]int, n)
    for i := range dist {
        dist[i] = make([]int, n)
        for j := range dist[i] {
            dist[i][j] = math.MaxInt32 / 2
        }
        dist[i][i] = 0
    }
    for _, edge := range edges {
        dist[edge[0]][edge[1]] = edge[2]
        dist[edge[1]][edge[0]] = edge[2]
    }
    
    for k:=0; k<n; k++ {
        for i:=0; i<n; i++ {
            for j:=0; j<n; j++ {
                if dist[i][j] > dist[i][k] + dist[k][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
                }
            }
        }
    }
    
    ansId := -1
    ansNum := n + 1
    for i:=0; i<n; i++ {
        num := 0
        for j:=0; j<n; j++ {
            if i ==j {
                continue
            }
            if dist[i][j] <= distanceThreshold {
                num++
            }
        }
        if num <= ansNum {
            ansNum = num
            ansId = i
        }
    }
    return ansId
}

