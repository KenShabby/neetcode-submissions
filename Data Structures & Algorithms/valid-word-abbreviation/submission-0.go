func validWordAbbreviation(word string, abbr string) bool {
	var i, j int
	jumpDist := 0

	for j < len(abbr){
		// test for leading zero
		if abbr[j] == '0' {
			return false
		}
		if isDigit(abbr[j]) {
			for j < len(abbr) && isDigit(abbr[j]) {
				jumpDist = jumpDist*10 + int(abbr[j]-'0')
				j++
			}
			i += jumpDist
			continue
		}
		jumpDist = 0
		if j < len(abbr) && i < len(word) && word[i] == abbr[j] {
			i++
			j++
		} else {
			return false
		}
	}

	if j == len(abbr) && i == len(word){
		return true
	} else {
		return false
	}
}

func isDigit (c byte) bool {
	if c >= '0' && c<= '9' {
		return true
	} else {
		return false
	}
}