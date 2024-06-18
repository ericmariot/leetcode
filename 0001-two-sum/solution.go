func twoSum(nums []int, target int) []int {
    resultIndices := make([]int, 2)

    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                resultIndices[0] = i
                resultIndices[1] = j
            }
        }
    }
    
    return resultIndices
}
