func findMaxAverage(nums []int, k int) float64 {
    maxSum := 0
    for i := 0; i < k; i++ {
        maxSum += nums[i]
    }

    for i := 0; i < len(nums)-k+1; i++ {
        curr := 0

        for j := range k {
            curr += nums[j+i]
        }

        if curr > maxSum {
            maxSum = curr
        }
    }

    return float64(maxSum) / float64(k)
}
