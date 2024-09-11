func characterReplacement(s string, k int) int {
    count := map[byte]int{}
    res := 0

    l := 0
    f := 0
    for r := range s {
        count[s[r]] = 1 + count[s[r]]
        f = max(f, count[s[r]])
        if (r - l + 1) - f > k {
            count[s[l]] -= 1
            l++
        }

        res = max(res, r - l + 1)
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
