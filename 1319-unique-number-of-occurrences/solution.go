func uniqueOccurrences(arr []int) bool {
    freq := map[int]int{}
    for _, n := range arr {
        freq[n]++
    }

    uniqueFreq := map[int]bool{}
    for _, n := range freq {
        if uniqueFreq[n] {
            return false
        }

        uniqueFreq[n] = true
    }

    return true
}
