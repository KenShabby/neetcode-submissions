func minOperations(logs []string) int {
	depth := 0
	for _, cmd := range logs{
		if cmd == "../" {
			if depth > 0 {
				depth--
				continue
			}
			continue
		}
		if cmd == "./" {
		} else {
			depth++
		}
	}
	return depth
}
