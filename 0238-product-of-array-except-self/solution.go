func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    pre := 1
    for i := range len(nums) {
        res[i] = pre
        pre *= nums[i]
    }

    pos := 1
    for i := len(nums)-1; i >= 0; i-- {
        res[i] *= pos
        pos *= nums[i]
    }

    return res
}

