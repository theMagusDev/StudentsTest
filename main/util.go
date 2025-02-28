package main

func createBoxToCardMap() map[int]int {
	cardsNumbers := getShuffledArray(50, 50)
	boxToCardMap := make(map[int]int)
	for boxNumber := 1; boxNumber <= 50; boxNumber++ {
		boxToCardMap[boxNumber] = cardsNumbers[boxNumber-1]
	}
	return boxToCardMap
}

func getShuffledArray(length, max int) []int {
	if length > max {
		panic("length cannot be greater than max")
	}

	numbers := make([]int, max)
	for i := range numbers {
		numbers[i] = i + 1
	}

	random.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	return numbers[:length]
}
