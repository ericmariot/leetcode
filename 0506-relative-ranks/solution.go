func findRelativeRanks(score []int) []string {
    answer := make([]string, len(score))
    aux := make([]int, len(score))

    for i := range aux {
        aux[i] = i
    }

    sort.Slice(aux, func(i, j int) bool {
        return score[aux[i]] > score[aux[j]]
    })

    for rank, idx := range aux {
        switch rank {
        case 0:
            answer[idx] = "Gold Medal"
        case 1:
            answer[idx] = "Silver Medal"
        case 2:
            answer[idx] = "Bronze Medal"
        default:
            answer[idx] = strconv.Itoa(rank + 1)
        }
    }

    return answer
}
