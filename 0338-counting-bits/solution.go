func countBits(n int) []int {
    var res []int

    
    for num := range n+1 {
        nRes := 0

        for num != 0 {
            if num & 1 == 1 {
                nRes++
            }
            num >>= 1
        }

        res = append(res, nRes)
    }

    return res
}
