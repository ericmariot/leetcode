func groupAnagrams(strs []string) [][]string {
    var res [][]string
    table := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int

        for _, c := range s {
            count[c - 'a']++
        }

        table[count] = append(table[count], s)
    }

    for _, g := range table {
        res = append(res, g)
    }

    return res
}
