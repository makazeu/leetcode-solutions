func reverseWords(s []byte)  {
    var r []byte
    e := len(s)
    for p := len(s) - 1; p >=0; p-- {
        if p == 0 || s[p-1] == ' ' {
            r = append(r, s[p:e]...)
            r = append(r, ' ')
            e = p - 1
        }
    }
    for i := range s {
        s[i] = r[i]
    }
}

