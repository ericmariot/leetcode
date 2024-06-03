func scoreOfString(s string) int {
    score := 0

    for i := 0; i < len(s)-1; i++ {
        score += diff(int(s[i]), int(s[i+1]))
    }

    return score
}

func diff(x int, y int) int {
    if x > y {
        return x - y
    }

    return y - x
}
