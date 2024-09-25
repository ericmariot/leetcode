func asteroidCollision(asteroids []int) []int {
    s := []int{asteroids[0]}

    for i := 1; i < len(asteroids); i++ {
        for len(s) > 0 && s[len(s)-1] > 0 && s[len(s)-1] < -asteroids[i] {
            s = s[:len(s)-1]
        }

        if len(s) == 0 || s[len(s)-1] < 0 || asteroids[i] > 0 {
            s = append(s, asteroids[i])
        }

        if s[len(s)-1] == -asteroids[i] {
            s = s[:len(s)-1]
        }
    }

    return s
}
