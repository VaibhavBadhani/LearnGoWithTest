package arraysandslices

// func Sum(numbers [5]int) int {
// 	var sum int
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

/****************************************************** Learning ******************************************************

	Array literals and type inference

	When declaring arrays, you have two options:
	1   b := [2]string{"Penn", "Teller"} // explicitly says array has length 2
		Here the type of b is [2]string.

	2   You can let the compiler count elements for you:
		b := [...]string{"Penn", "Teller"}
		Here [...] tells Go: count the number of elements and infer the size.
		The type is still [2]string.


	Arrays in Go are fixed-size values.
	Assigning or passing them copies the whole array.
	To share the same underlying data, use pointers or, more commonly in Go, slices.
	Arrays are rarely used directly; slices are preferred because they’re lightweight references to arrays with dynamic length handling.




****************************************************** END ******************************************************/
