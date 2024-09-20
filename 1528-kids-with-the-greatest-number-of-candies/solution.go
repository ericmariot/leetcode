func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := []bool{}
    h := 0

    for _, v := range candies {
        if v > h {
            h = v
        }
    }

    for _, v := range candies {
        if v + extraCandies >= h {
            res = append(res, true)
            continue
        }

        res = append(res, false)
    }
    
    return res
}
