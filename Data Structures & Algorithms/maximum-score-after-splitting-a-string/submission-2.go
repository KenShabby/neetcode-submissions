func maxScore(s string) int {
    ones := strings.Count(s, "1")
    zeros := 0
    best := 0

    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            zeros++
        } else {
            ones--
        }
        if zeros+ones > best {
            best = zeros + ones
        }
    }
    return best
}