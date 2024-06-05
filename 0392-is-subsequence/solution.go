func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
		return true
	}

	sPos := 0

	for _, char := range t {
		if string(char) == string(s[sPos]) {
			sPos++
			if sPos == len(s) {
				return true
			}
		}
	}

	if sPos < len(s) {
		return false
	}

	return true
}
