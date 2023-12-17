package reducefn

// Code copied from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/revisiting-arrays-and-slices-with-generics

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	return Reduce[int](numbers, func(acc, i2 int) int {
		return acc + i2
	}, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	return Reduce[[]int](numbersToSum, func(acc, i2 []int) []int {
		if len(i2) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(i2[1:]))
		}
	}, []int{})
}

// Again use generics to DRY up our code
// create a reduce function

func Reduce[T any](coll []T, accum func(T, T) T, initVal T) T {
	res := initVal
	for _, val := range coll {
		res = accum(res, val)
	}
	return res
}
