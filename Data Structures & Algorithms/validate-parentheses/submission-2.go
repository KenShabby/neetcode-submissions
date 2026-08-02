func isValid(s string) bool {
    stack := []rune{}

    for _, char := range s {
        if char == '(' || char == '[' || char == '{' {
            stack = append(stack, char)
        } 
        if char == ')' || char == ']' || char == '}' {
            if len(stack) > 0 {
                topOfStack := stack[len(stack)-1]
                if topOfStack == '(' && char == ')' {
                    stack = stack[:len(stack)-1]
                } else if topOfStack == '[' && char == ']' {
                    stack = stack[:len(stack)-1]
                } else if topOfStack == '{' && char == '}' {
                    stack = stack[:len(stack)-1]
                } else {
                return false
                }
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}

