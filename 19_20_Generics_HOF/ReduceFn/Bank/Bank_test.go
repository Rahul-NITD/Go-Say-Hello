package bank_test

import (
	bank "GoSayHello/19_20_Generics_HOF/ReduceFn/Bank"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v != %v", got, want)
	}
}

func TestBank(t *testing.T) {
	transactions := []bank.Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual[int](t, bank.BalanceFor(transactions, "Riya"), 100)
	AssertEqual[int](t, bank.BalanceFor(transactions, "Chris"), -75)
	AssertEqual[int](t, bank.BalanceFor(transactions, "Adil"), -25)

}
