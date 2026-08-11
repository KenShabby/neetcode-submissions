func countSeniors(details []string) int {
   nSeniors := 0

   for _, passenger := range details {
	tens, _ := strconv.Atoi(string(passenger[11])) 
	tens *= 10
	ones, _ := strconv.Atoi(string(passenger[12]))
	age := tens + ones
	if age > 60 {
		nSeniors++
	}
   } 
   return nSeniors
}