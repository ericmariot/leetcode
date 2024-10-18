func twoSum(nums []int, target int) []int {
    h := map[int]int{}

    for i, num := range nums {
        diff := target - num

        if idx, found := h[diff]; found {
            return []int{i, idx}
        }
        
        h[num] = i
    }

    return nil
}
