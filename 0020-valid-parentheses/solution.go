func isValid(s string) bool {
    if len(s) < 2 {
        return false
    }

    var stack []rune
    closeToOpen := map[string]string{")": "(", "]": "[", "}": "{"}
    
    for _, r := range s {
        val, ok := closeToOpen[string(r)]
        if ok {
            if len(stack) > 0 && string(stack[len(stack)-1]) == val {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else {
            stack = append(stack, r)
        }
    }

    if len(stack) != 0 {
        return false
    }

    return true
}
