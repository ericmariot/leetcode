func maxOperations(nums []int, k int) int {
    m := map[int]int{}

    for _, n := range nums {
        m[n]++
    }

    res := 0

    for _, n := range nums {
        if n >= k {
            continue
        }

        t := k - n

        if t != n && m[t] > 0 && m[n] > 0 {
            res++
            m[t]--
            m[n]--
        }

        if t == n && m[t] > 1 && m[n] > 1 {
            res++
            m[t]--
            m[n]--
        }

        if m[t] == 0 {
            delete(m, t)
        }

        if m[n] == 0 {
            delete(m, n)
        }
    }

    return res
}
