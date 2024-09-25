func tribonacci(num int) int {
    if num == 0 {
        return 0
    }

    if num == 1 {
        return 1
    }

    x, y, z := 0, 1, 1
    for i := 2; i < num; i++ {
        n := x + y + z
        x, y, z = y, z, n
    }

    return z
}
