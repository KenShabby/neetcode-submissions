func lengthOfLastWord(s string) int {
	str := strings.TrimSpace(s)	
	if len(str) == 0 {
		return 0
	}
	count := 0

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' '{
			fmt.Printf("str[%d]: %c\n", i, str[i])
			return count
		} else {
			fmt.Printf("str[%d]: %c\n", i, str[i])
			count++
		}
	}
	return count
}