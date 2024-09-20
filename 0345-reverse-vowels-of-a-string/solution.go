func reverseVowels(s string) string {
    left, right := 0, len(s)-1
    vowels := "aeiouAEIOU"
    runes := []rune(s)

    for left < right {
        for left < right && !strings.Contains(vowels, string(runes[left])) {
            left++
        }
        
        for left < right && !strings.Contains(vowels, string(runes[right])) {
            right--
        }

        if left < right {
            runes[left], runes[right] = runes[right], runes[left]
            left++
            right--
        }
    }
    
    return string(runes)
}

