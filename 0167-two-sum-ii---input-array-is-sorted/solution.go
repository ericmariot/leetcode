func twoSum(nums []int, t int) []int {
    b := 0
    e := len(nums)-1

    for e >= b {
        switch {
            case nums[b] + nums[e] > t:
                e--
            case nums[b] + nums[e] < t:
                b++
            default:
                return []int{b+1, e+1}
        }
    }

    return []int{}
}
