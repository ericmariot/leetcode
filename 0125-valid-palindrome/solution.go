func isPalindrome(s string) bool {
    reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}

    cleaned := reg.ReplaceAllString(strings.ToLower(s), "")
    
    for i, j := 0, len(cleaned)-1; i < len(cleaned); i, j = i+1, j-1 {
        if cleaned[i] != cleaned[j] {
            return false
        }
    }


    return true
}
