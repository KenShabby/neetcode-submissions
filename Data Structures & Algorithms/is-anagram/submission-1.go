import "maps"

func isAnagram(s string, t string) bool {
	hashTable1 := make(map[rune]int)
	hashTable2 := make(map[rune]int)

	for _, char := range s {
		hashTable1[char]++
	}
	for _, char := range t {
		hashTable2[char]++
	}

	if maps.Equal(hashTable1, hashTable2) {
		return true
	} else {
		return false
	}
}
