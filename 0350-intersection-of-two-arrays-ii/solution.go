func intersect(nums1 []int, nums2 []int) []int {
    var small, big, result []int
    m := make(map[int]int)

    if len(nums1) > len(nums2) {
        big = nums1
        small = nums2
    } else {
        big = nums2
        small = nums1
    }

    for _, num := range small {
        m[num]++
    }

    for _, num := range big {
        _, ok := m[num]
        if ok {
            if m[num] > 0 {
                result = append(result, num)
                m[num]--
            }
        }
    }


    return result
}
