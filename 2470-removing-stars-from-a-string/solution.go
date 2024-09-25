func removeStars(s string) string {
    x := []string{}

    for _, v := range s {
        x = append(x, string(v))
        if string(v) == "*" {
            x = x[:len(x)-2]
        }
    }
    
    return strings.Join(x[:], "")
}
