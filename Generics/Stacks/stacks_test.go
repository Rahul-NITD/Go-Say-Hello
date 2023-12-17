package stacks_test

import (
	stacks "GoSayHello/Generics/Stacks"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v != %v", got, want)
	}
}

func TestStack(t *testing.T) {
	t.Run("For Integers", func(t *testing.T) {
		st := new(stacks.StackOfInt)
		AssertEqual(t, st.IsEmpty(), true)
		st.Push(5)
		AssertEqual(t, st.IsEmpty(), false)
		val, err := st.Pop()
		AssertEqual(t, val, 5)
		AssertEqual(t, err, nil)
		AssertEqual(t, st.IsEmpty(), true)
	})
}
