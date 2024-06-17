func isPerfectSquare(num int) bool {
    left := 0
    right := num + 1

    for left < right {
        mid := left + (right - left) / 2

        if mid * mid >= num {
            right = mid
            continue
        }

        left = mid + 1
    }

    return left * left == num
}
