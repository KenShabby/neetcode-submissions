func calPoints(operations []string) int {
	var scoreStack []int
	
	for _, op := range operations {
		switch op {
		case "+":
			n := len(scoreStack)
			scoreStack = append(scoreStack, scoreStack[n-1] + scoreStack[n-2])
		case "D":
			scoreStack = append(scoreStack, scoreStack[len(scoreStack)-1]*2)
		case "C":
			scoreStack = scoreStack[:len(scoreStack)-1]
		default:
			n, _ := strconv.Atoi(op)
			scoreStack = append(scoreStack, n)
		}
	}

	sum := 0
	for _, val := range scoreStack {
		sum += val
	}
	return sum
}