
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}
	if len(s) > len(t){
		return false
	}

	for i := 0; i<len(s); {
		for j := 0; j < len(t) && i < len(s); {
			if s[i] == t[j] {
				i++
				j++
			} else {
				j++
			}
		}
		if i == len(s){
			return true
		} else { 
			return false
		}
	}
	return false
}
