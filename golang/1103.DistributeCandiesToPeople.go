func distributeCandies(candies int, num_people int) []int {
    now := 1
    ans := make([]int, num_people)
    for {
        for i:=0;i<num_people;i++ {
            if(candies > now) {
                ans[i] += now
                candies -= now
                now++
            } else {
                ans[i] += candies
                candies = 0
                break
            }
        }
        if candies == 0 {
            break
        }
    }
    return ans
}

