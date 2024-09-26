func majorityElement(nums []int) int {
    element := 0
    counter := 0
    
    for _, n := range nums {
        if counter == 0 {
            element = n
            counter = 1
        } else if element == n {
            counter += 1
        } else {
            counter -= 1
        }
    }

    return element
}
