func isIsomorphic(s string, t string) bool {
	var sMap = make(map[byte]byte)
	var tMap = make(map[byte]byte)

	for i := 0; i<len(s); i++ {
		c1, c2 := s[i], t[i]

		if v, ok := sMap[c1]; ok {
			if v != c2 {
				return false
			} 
		} else {
				sMap[c1] = c2
		}

		if v, ok := tMap[c2]; ok {
			if v != c1 {
				return false
			} 
		} else {
				tMap[c2] = c1
		}
	}
	return true
}