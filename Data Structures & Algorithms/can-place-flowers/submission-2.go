func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i, bed := range flowerbed {
		if bed == 0 {
			leftOK := i == 0 || flowerbed[i-1] == 0
			rightOK := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if leftOK && rightOK {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}