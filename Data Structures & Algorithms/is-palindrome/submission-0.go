func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		l := rune(s[left])
		r := rune(s[right])

		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
			continue
		}

		if unicode.ToLower(l) != unicode.ToLower(r) {
			return false
		}
		left++
		right--
	}
	return true
}
