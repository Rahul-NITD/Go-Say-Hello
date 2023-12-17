package reducefn_test

import (
	reducefn "GoSayHello/Generics/ReduceFn"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v != %v", got, want)
	}
}

func TestSum(t *testing.T) {
	AssertEqual(t, reducefn.Sum([]int{1, 2, 3, 4, 5}), 15)
	AssertEqual(t, reducefn.SumAllTails([]int{1, 2, 3, 4, 5})[0], 14)
}
