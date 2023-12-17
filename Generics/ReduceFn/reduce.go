package reducefn

// Code copied from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/revisiting-arrays-and-slices-with-generics

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	return Reduce[int](numbers, func(i1, i2 int) int {
		return i1 + i2
	})
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
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

// Again use generics to DRY up our code
// create a reduce function

func Reduce[T any](coll []T, accum func(T, T) T) T {
	var res T
	for _, val := range coll {
		res = accum(res, val)
	}
	return res
}
