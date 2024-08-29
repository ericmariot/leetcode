func longestConsecutive(nums []int) int {
    set := map[int]bool{}

    for _, n := range nums {
        set[n] = true
    }

    count := 0

    for _, n := range nums {
        if set[n-1] {
            continue
        }

        l := 1
        tmp := n + 1
        for set[tmp] {
            l++
            tmp++
        }

        if l > count {
            count = l
        }
    }

    return count
}
