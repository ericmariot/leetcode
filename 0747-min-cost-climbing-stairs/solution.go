func minCostClimbingStairs(cost []int) int {
    step1 := 0
    step2 := 0

    for i := 2; i <= len(cost); i++ {
        f := step1 + cost[i-1]
        s := step2 + cost[i-2]
        step2 = step1
        step1 = min(f, s)
    }

    return step1
}
