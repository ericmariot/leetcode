func findDifference(nums1 []int, nums2 []int) [][]int {
    res := make([][]int, 2)
    mNums1 := map[int]bool{}
    mNums2 := map[int]bool{}

    for _, n := range nums1 {
        mNums1[n] = true
    }

    for _, n := range nums2 {
        mNums2[n] = true
    }

    for k := range mNums1 {
        if !mNums2[k] {
            res[0] = append(res[0], k)
        }
    }

    for k := range mNums2 {
        if !mNums1[k] {
            res[1] = append(res[1], k)
        }
    }

    return res
}
