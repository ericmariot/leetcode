func pivotIndex(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }

    leftSum := 0
    rightSum := sum

    for i, v := range nums {
        rightSum -= v

        if leftSum == rightSum {
            return i
        }

        leftSum += v
    }

    return -1
}
