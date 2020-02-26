func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    x := 0
    var t int 
    var result string
    for num > 0 {
        t = num % 1000
        num /= 1000
        if t > 0 {
            result = getBelowThousand(t) + " " + sec[x] + " " + result
        }
        x++
    }

    return strings.TrimSpace(result)
}

var sec = []string{"", "Thousand", "Million", "Billion"}
var tens = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var decades = []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var ones = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func getBelowThousand(num int) string {
    if num == 0 {
        return ""
    }
    var result string
    if num >= 100 {
        t := num / 100
        result = ones[t] + " Hundred"
        num %= 100
    }
    if num >= 20 {
        t := num / 10
        num %= 10
        result = result + " " + decades[t-2]
    } else if num >= 10 {
        result = result + " " + tens[num-10]
        return strings.TrimSpace(result)
    } 
    result = result + " " + ones[num]
    return strings.TrimSpace(result)
}

