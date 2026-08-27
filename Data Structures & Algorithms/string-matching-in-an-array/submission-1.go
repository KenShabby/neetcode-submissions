func stringMatching(words []string) []string {
	var res []string

	for i, word := range words {
		found := false
		for j, other := range words {
			if i != j && strings.Contains(other, word) {
				found = true
				break
			}
		}
		if found == true {
			res = append(res, word)
		}
	}
	return res
}