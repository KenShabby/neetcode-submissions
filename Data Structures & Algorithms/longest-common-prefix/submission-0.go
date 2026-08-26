import "slices"

func longestCommonPrefix(strs []string) string {
   res := ""
   if len(strs) == 0 {
	return res
   }
   if len(strs) == 1 {
	return strs[0]
   }

   slices.Sort(strs)

	for i := 0; i < len(strs[0]); i++ {
		if strs[0][i] == strs[len(strs)-1][i] {
			res += string(strs[0][i])
		} else {
			break
		}
	}
	return res
}
