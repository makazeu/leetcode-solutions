func angleClock(hour int, minutes int) float64 {
    hour %= 12
    ans := float64(minutes) / 60 - float64(hour * 60 + minutes) / float64(60*12)
    ans *= 360
    
    if ans < 0 {
        ans = -ans
    }
    if ans > 180 {
        ans = 360-ans
    }
    return ans
}

