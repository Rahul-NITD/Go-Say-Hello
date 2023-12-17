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

func TestReduce(t *testing.T) {
	t.Run("Multiply Elements", func(t *testing.T) {
		AssertEqual[int](t, reducefn.Reduce[int](
			[]int{1, 2, 3, 5}, func(i1, i2 int) int {
				return i1 * i2
			},
			1,
		), 30)
	})
	t.Run("Concat Strongs using Reduce", func(t *testing.T) {
		res := reducefn.Reduce[string](
			[]string{"1", "2", "3"},
			func(s1, s2 string) string {
				return s1 + s2
			}, "",
		)
		AssertEqual[string](t, res, "123")
	})

	t.Run("Reduce Structs", func(t *testing.T) {

		type AppleBasket struct {
			apples   int
			polished bool
		}

		applesBaskets := []AppleBasket{
			{5, true},
			{4, false},
			{8, true},
		}

		t.Run("Sum All Apples", func(t *testing.T) {
			res := reducefn.Reduce[AppleBasket, int](
				applesBaskets,
				func(valTillNow int, Curr AppleBasket) int {
					return valTillNow + Curr.apples
				},
				0,
			)
			AssertEqual[int](t, res, 17)
		})

		t.Run("Sum only polished Apples", func(t *testing.T) {
			// res := reducefn.Reduce[AppleBasket](
			// 	applesBaskets,
			// 	func(ab1, ab2 AppleBasket) AppleBasket {
			// 		if !ab2.polished {
			// 			return ab1
			// 		}
			// 		return AppleBasket{
			// 			ab1.apples + ab2.apples,
			// 			true,
			// 		}
			// 	}, AppleBasket{0, true},
			// )
			// Now we can do this instead
			res := reducefn.Reduce[AppleBasket, int](
				applesBaskets,
				func(valTillNow int, curr AppleBasket) int {
					if curr.polished {
						return valTillNow + curr.apples
					}
					return valTillNow
				},
				0,
			)
			AssertEqual[int](t, res, 13)
		})
	})
}
