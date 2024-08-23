func rob(nums []int) int {
    first, second := 0, 0
    // [first, second, n, n+1, ...]

    for _, num := range nums {
        tmp := 0
        if first + num > second {
            tmp = first + num
        } else {
            tmp = second
        }

        first = second
        second = tmp
    }

    return second
}
