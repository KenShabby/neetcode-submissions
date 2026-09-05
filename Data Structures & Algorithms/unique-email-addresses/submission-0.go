func numUniqueEmails(emails []string) int {

	addressHash := make(map[string]any)

	for _, email := range emails{
		local, hostname, _ := strings.Cut(email, "@")
		
		// Handle local name portion
		local, _, _ = strings.Cut(local, "+") // Truncate after '+'
		var bytes strings.Builder
		for _, ch := range local {
			if ch == '.' {
				continue
			}
			bytes.WriteByte(byte(ch))
		}
		local = bytes.String()

		// Handle domain portion
		var b strings.Builder
		for _, ch := range hostname {
			b.WriteByte(byte(ch))
		}
		hostname = b.String()
		addressHash[local + "@" + hostname] = true
		
	}

	return len(addressHash)
}