func isValid(s string) bool {
    if len(s) < 2 {
        return false
    }

    var stack []rune

    for _, r := range s {
        if r == 40 || r == 91 || r == 123 {
            stack = append(stack, r)
            continue
        }

        if len(stack) == 0 && (r == 41 || r == 93 || r == 125) {
            return false
        }

        if len(stack) > 0 {
            if r == 41 {
                if stack[len(stack)-1] == 40  {
                    stack = stack[:len(stack)-1]
                    continue
                }
                return false
            }

            if r == 93 {
                if stack[len(stack)-1] == 91 {
                    stack = stack[:len(stack)-1]
                    continue
                }
                return false
            }

            if r == 125 {
                if stack[len(stack)-1] == 123 {
                    stack = stack[:len(stack)-1]
                    continue
                }
                return false
            }
        }
    }

    if len(stack) != 0 {
        return false
    }

    return true
}
