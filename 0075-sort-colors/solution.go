func sortColors(nums []int)  {
    red, white := 0, 0
    for _, color := range(nums) {
        if color == 0 {
            red++
        }
        if color == 1 {
            white++
        }
    }

    for i := 0; i < red; i++ {
        nums[i] = 0
    }

    for i := red; i < red+white; i++ {
        nums[i] = 1
    }

    for i := red+white; i < len(nums); i++ {
        nums[i] = 2
    }
}
