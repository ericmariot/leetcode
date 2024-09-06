func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    var res [][]int

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        j := i + 1
        k := len(nums) -1

        for j < k {
            s := nums[j] + nums[k] + nums[i]
            if s == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})

                k--

                for j < k && nums[k] == nums[k+1] {
                    k--
                }
            } else if s > 0 {
                k--
            } else {
                j++
            }
        }
    }


    return res
}
