func twoSum(numbers []int, target int) []int {
    start := 0
    end := len(numbers)-1

    for end >= start {
        if numbers[start] + numbers[end] == target {
            return []int{start+1, end+1}
        }

        if numbers[start] + numbers[end] > target {
            end--
        }

        if numbers[start] + numbers[end] < target {
            start++
        }
    }

    return []int{}
}
