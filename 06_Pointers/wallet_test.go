package shitbank

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want ShitCoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted Error Didn't get it")
		}
		if got != want {
			t.Errorf("Got '%q' Want '%q'", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("Did not want an error here")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(ShitCoin(10))

		want := ShitCoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(ShitCoin(10))
		err := wallet.Withdraw(ShitCoin(5))
		want := ShitCoin(5)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("aukat k bahr withdraw", func(t *testing.T) {
		wallet := Wallet{ShitCoin(10)}
		err := wallet.Withdraw(ShitCoin(100))
		assertError(t, err, InsufficientFundsError)
		assertBalance(t, wallet, ShitCoin(10))
	})

}
