package main

func SumAllTails(numbersToSumTail ...[]int) []int {
	var sums []int

	for _, number := range numbersToSumTail {
		var tail = []int{}
		if len(number) > 1 {
			tail = number[1:]
		}
		sums = append(sums, Sum(tail))
	}
	return sums
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
