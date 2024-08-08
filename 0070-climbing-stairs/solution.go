func climbStairs(n int) int {
    last, secondLast := 1, 0
    for i := 1; i <= n; i++ {
        last, secondLast = secondLast + last, last
    }
    return last
}
