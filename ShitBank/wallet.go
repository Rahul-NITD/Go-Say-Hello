package shitbank

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("Insufficient ShitCoins, Withdraw Cancelled")

// Our Bank only runs on ShitCoins
type ShitCoin int

type Wallet struct {
	balance ShitCoin
}

func (wallet *Wallet) Deposit(amt ShitCoin) {
	wallet.balance += amt
}

func (wallet *Wallet) Withdraw(amt ShitCoin) error {
	if amt > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= amt
	return nil
}

func (wallet *Wallet) Balance() ShitCoin {
	return wallet.balance
}

type Stringer interface {
	String() string
}

func (sc ShitCoin) String() string {
	return fmt.Sprintf("%d STC", sc)
}
