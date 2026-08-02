func getConcatenation(nums []int) []int {
   var ans []int 
   ans = nums
   for _, num := range nums {
	ans = append(ans, num)
   }

   return ans
}
