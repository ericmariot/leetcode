func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := make(map[rune]int, len(s))
    for _, l := range s {
        m[l]++
    }

    n := make(map[rune]int, len(t))
    for _, l := range t {
        n[l]++
    }

    if !maps.Equal(m, n) {
        return false
    }

    return true
}
