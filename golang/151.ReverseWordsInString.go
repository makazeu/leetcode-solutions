func reverseWords(s string) string {
    var result string
    for _, word := range strings.Split(s, " ") {
        if len(strings.TrimSpace(word)) == 0 {
            continue
        }
        result = word + " " + result
    }
    if len(result) > 0 {
        result = result[:len(result) - 1]
    }
    return result
}

