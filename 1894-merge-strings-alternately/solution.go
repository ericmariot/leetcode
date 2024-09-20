func mergeAlternately(word1 string, word2 string) string {
    maxLen := maxLength(word1, word2)
    res := []string{}

    for i := 0; i < maxLen; i++ {
        if i < len(word1) {
            res = append(res, string(word1[i]))
        }
        if i < len(word2) {
            res = append(res, string(word2[i]))
        }
    }

    return strings.Join(res, "")
}

func maxLength(w1, w2 string) int {
    if len(w1) > len(w2) {
        return len(w1)
    }

    return len(w2)
}
