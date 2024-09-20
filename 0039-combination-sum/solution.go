func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    curr := make([]int, 0)

    var bt func(index int, sum int, curr []int)
    bt = func(index int, sum int, curr []int) {
        if sum == target {
            res = append(res, append([]int{}, curr...))
            return
        }

        if sum > target {
            return
        }

        for i := index; i < len(candidates); i++ {
            curr = append(curr, candidates[i])
            bt(i, sum+candidates[i], curr)
            curr = curr[:len(curr)-1]
        }
    }
    bt(0, 0, curr)
    return res
}
