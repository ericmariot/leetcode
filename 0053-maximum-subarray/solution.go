func maxSubArray(nums []int) int {
    maxSub := nums[0]
    curr := 0

    for _, n := range nums {
        if curr < 0 {
            curr = 0
        }
        curr += n
        maxSub = max(maxSub, curr)
    }

    return maxSub
}

func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    }

    return num2
}
