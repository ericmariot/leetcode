func twoSum(nums []int, target int) []int {
    m := map[int]int{}

    for i, num := range nums {
        e, ok := m[target-num]
        if ok {
            return []int{e, i}
        }

        m[num] = i
    }

    return []int{}
}
