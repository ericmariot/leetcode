func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
    count := make([][]int, len(nums)+1)
    res := []int{}

    for _, n := range nums {
        if c, ok := m[n]; ok {
            m[n] = c + 1
        } else {
            m[n] = 1
        }
    }

    for n, c := range m {
        count[c] = append(count[c], n)
    }

    for i := len(count) - 1; i > 0; i-- {
        res = append(res, count[i]...)
        if len(res) == k {
            break
        }
    }

    return res
}
