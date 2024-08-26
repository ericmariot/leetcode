func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)

    for _, n := range nums {
        m[n]++
    }

    pairs := make([][2]interface{}, 0, len(m))
    for k, v := range m {
        pairs = append(pairs, [2]interface{}{k, v})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1].(int) > pairs[j][1].(int)
    })

    res := make([]int, k)
    for i, p := range pairs {
        if i > k-1 {
            break
        }
        res[i] = p[0].(int)
    }

    return res
}
