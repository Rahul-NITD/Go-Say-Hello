package choochoo

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test for 5 elements", func(t *testing.T) {
		numbers := []int{1, 4, 2, 5, 6}
		got := Sum(numbers)
		want := 18
		assertHelper(t, got, want)

	})
}

func TestSumAll(t *testing.T) {
	t.Run("test Sum All function for inputs {1,2,1} and {5,1,2}", func(t *testing.T) {
		a := []int{1, 2, 1}
		b := []int{5, 1, 2}
		got := SumAll(a, b)
		want := []int{4, 8}
		assertHelper(t, got, want)
	})
	t.Run("test with 5 inputs", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		c := []int{1, 3, 5}
		d := []int{2, 4, 6}
		e := []int{1, 4, 9}
		got := SumAll(a, b, c, d, e)
		want := []int{6, 15, 9, 12, 14}
		assertHelper(t, got, want)
	})
}

func assertHelper(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

// func assertHelper(t testing.TB, got, want) {
// 	t.Helper()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v for '%v' and '%v'", got, want, a, b)
// 	}
// }
