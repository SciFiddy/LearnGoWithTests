package main

func sum(numbers []int) int {

	output := 0

	for _, num := range numbers {
		output += num
	}

	return output
}

func SumAll(numbersToSum ...[]int) []int {

	sums := make([]int, len(numbersToSum))

	return sums
}
