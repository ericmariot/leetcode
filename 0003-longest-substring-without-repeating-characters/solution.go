func lengthOfLongestSubstring(s string) int {
    m := make([]int, 128)
    res := 0

    for i, j := 0, 0; i < len(s); i++ {
        m[s[i]]++
        for m[s[i]] > 1 {
            m[s[j]]--
            j++
        }
        if i-j+1 > res {
            res = i - j + 1
        }
    }

    return res
}
