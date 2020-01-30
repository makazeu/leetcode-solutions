func findRestaurant(list1 []string, list2 []string) []string {
    hash := make(map[string]int)
    for i, str := range list1 {
        hash[str] = i
    }
    
    leastIndex := math.MaxInt32
    var result []string
    
    for i, str := range list2 {
        if idx, exists := hash[str]; exists {
            if idx + i < leastIndex {
                leastIndex = idx + i
                result = append([]string{}, str)
            } else if idx + i == leastIndex {
                result = append(result, str)
            }
        }
    }
    
    return result
}

