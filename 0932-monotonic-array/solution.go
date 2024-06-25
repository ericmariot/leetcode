func isMonotonic(nums []int) bool {
    isIncreasing, isDecreasing := true, true

    for i := 0; i < len(nums)-1; i++ {
        num := nums[i]
        next := nums[i+1]

        if next > num {
            isDecreasing = false
        }

        if next < num {
            isIncreasing = false
        }
    }
    
    return isIncreasing || isDecreasing
}
