func largestAltitude(gain []int) int {
    res := 0
    curr := 0

    for _, v := range gain {
        curr = curr + v
        if curr > res {
            res = curr
        }
    }

    return res
}
