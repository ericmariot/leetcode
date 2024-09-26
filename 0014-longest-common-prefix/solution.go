func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    res := strs[0]
    for i := range res {
        for _, w := range strs[1:] {
            if i == len(w) || w[i] != res[i] {
                return res[:i]
            }
        }
    }

    return res
}
