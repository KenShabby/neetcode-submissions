func hasDuplicate(nums []int) bool {
   var hashTable = make(map[int]bool)

   for _, num := range nums {
	if hashTable[num] == true {
		return true
	}
	hashTable[num] = true
   } 

   return false

}
